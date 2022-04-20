package usecase

import (
	"errors"
	"time"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/Almazatun/golephant/pkg/http/presentation/_type"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	"github.com/Almazatun/golephant/pkg/util"
	"gopkg.in/go-playground/validator.v9"
)

type companyUseCase struct {
	companyRepo repository.CompanyRepo
}

type CompanyUseCase interface {
	RegisterCompany(registerCompanyInput input.RegisterCompanyInput) (companyDB *entity.Company, err error)
	LogIn(logInCompanyInput input.LogInCompanyInput) (res *_type.ResLogIn[entity.Company], err error)
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

func (uc *companyUseCase) LogIn(logInCompanyInput input.LogInCompanyInput) (res *_type.ResLogIn[entity.Company], err error) {

	company, err := uc.companyRepo.GetByEmail(logInCompanyInput.Email)

	if err != nil {
		return nil, err
	}

	isCorrectPassword := util.CheckPassword(logInCompanyInput.Password, company.Password)

	if !isCorrectPassword {
		newErr := errors.New(error_message.INCCORECT_PASSWORD)
		return nil, newErr
	}

	generateJWT, err := jwt_gl.GenerateJWTStr(company.Email)

	if err != nil {
		return nil, err
	}

	return &_type.ResLogIn[entity.Company]{
		Token:             generateJWT.Token,
		ExperationTimeJWT: generateJWT.ExperationTime,
		LogInEntityData:   *company}, nil
}
