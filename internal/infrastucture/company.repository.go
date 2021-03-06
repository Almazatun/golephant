package repository

import (
	"errors"

	entity "github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

type CompanyRepo interface {
	Create(company entity.Company) (companyDB *entity.Company, err error)
	GetByEmail(email string) (companyDB *entity.Company, err error)
	GetByPhone(phone string) (companyDB *entity.Company, err error)
	GetById(companyId string) (companyDB *entity.Company, err error)
	Save(company entity.Company) (companyDB *entity.Company, err error)
}

func NewCompanyRepo(db *gorm.DB) CompanyRepo {
	return &companyRepository{
		db: db,
	}
}

func (r *companyRepository) Create(company entity.Company) (companyDB *entity.Company, err error) {
	result := r.db.Create(&company)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) GetByEmail(email string) (companyDB *entity.Company, err error) {
	var company entity.Company

	result := r.db.
		Preload("CompanyAddresses").
		First(&company, "email = ?", email)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.COMPANY_NOT_FOUND)

		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) GetById(companyId string) (companyDB *entity.Company, err error) {
	var company entity.Company

	result := r.db.
		Preload("CompanyAddresses").
		First(&company, "company_id = ?", companyId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.COMPANY_NOT_FOUND)

		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) GetByPhone(phone string) (companyDB *entity.Company, err error) {
	var company entity.Company

	result := r.db.
		Preload("CompanyAddresses").
		First(&company, "phone = ?", phone)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.COMPANY_NOT_FOUND)

		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) Save(company entity.Company) (companyDB *entity.Company, err error) {
	result := r.db.
		Preload("CompanyAddresses").
		Save(&company)

	e := result.Error

	if e != nil {
		return nil, e
	}

	return &company, nil
}
