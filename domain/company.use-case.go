package usecase

import (
	"errors"
	"time"

	common "github.com/Almazatun/golephant/common"
	error_message "github.com/Almazatun/golephant/common/error-message"
	repository "github.com/Almazatun/golephant/infrastucture"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
	types "github.com/Almazatun/golephant/presentation/types"
	"github.com/Almazatun/golephant/util"
	"gopkg.in/go-playground/validator.v9"
)

type companyUseCase struct {
	companyRepo repository.CompanyRepo
}

type CompanyUseCase interface {
	RegisterCompany(registerCompanyInput input.RegisterCompanyInput) (companyDB *entity.Company, err error)
	LogIn(logInCompanyInput input.LogInCompanyInput) (res *types.ResLogIn[entity.Company], err error)
}

func NewCompanyUseCase(companyRepo repository.CompanyRepo) CompanyUseCase {
	return &companyUseCase{
		companyRepo: companyRepo,
	}
}

type ResLogInCompany struct {
	Token             string
	ExperationTimeJWT time.Time
}

func (uc *companyUseCase) RegisterCompany(registerCompanyInput input.RegisterCompanyInput) (comapanyDB *entity.Company, err error) {
	// Validate register company input
	v := validator.New()
	e := v.Struct(registerCompanyInput)

	if e != nil {
		return nil, e
	}

	// Delete white space
	password := util.TrimWhiteSpace(registerCompanyInput.Password)
	email := util.TrimWhiteSpace(registerCompanyInput.Email)

	// Hashing company password
	hashedPassword, err := util.HashPassword(password)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	registerCompany := &entity.Company{
		Email:        email,
		Password:     hashedPassword,
		CreationTime: now,
		UpdateTime:   now,
	}

	res, err := uc.companyRepo.Create(*registerCompany)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *companyUseCase) LogIn(logInCompanyInput input.LogInCompanyInput) (res *types.ResLogIn[entity.Company], err error) {

	company, err := uc.companyRepo.GetByEmail(logInCompanyInput.Email)

	if err != nil {
		return nil, err
	}

	isCorrectPassword := util.CheckPassword(logInCompanyInput.Password, company.Password)

	if !isCorrectPassword {
		newErr := errors.New(error_message.INCCORECT_PASSWORD)
		return nil, newErr
	}

	generateJWT, err := common.GenerateJWTStr(company.Email)

	if err != nil {
		return nil, err
	}

	return &types.ResLogIn[entity.Company]{
		Token:             generateJWT.Token,
		ExperationTimeJWT: generateJWT.ExperationTime,
		LogInEntityData:   *company}, nil
}
