package repository

import (
	"errors"

	entity "github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type companyAddressRepository struct {
	db *gorm.DB
}

type CompanyAddressRepo interface {
	GetById(companyId string) (companyAddressDB *entity.CompanyAddress, err error)
	GetByTitle(title string) (companyAddressDB *entity.CompanyAddress, err error)
	Save(companyAddress entity.CompanyAddress) (companyAddressDB *entity.CompanyAddress, err error)
	DeleteById(companyId string) (str string, err error)
}

func NewCompanyAddressRepo(db *gorm.DB) CompanyAddressRepo {
	return &companyAddressRepository{
		db: db,
	}
}

func (r *companyAddressRepository) GetById(
	companyAddressId string,
) (companyAddressDB *entity.CompanyAddress, err error) {
	var companyAddress entity.CompanyAddress

	result := r.db.First(&companyAddress, "company_address_id = ?", companyAddressId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.COMPANY_ADDRESS_NOT_FOUND)

		return nil, err
	}

	return &companyAddress, nil
}

func (r *companyAddressRepository) GetByTitle(
	title string,
) (companyAddressDB *entity.CompanyAddress, err error) {
	var companyAddress entity.CompanyAddress

	result := r.db.First(&companyAddress, "title = ?", title)

	dbErr := result.Error

	if dbErr != nil {
		return nil, err
	}

	return &companyAddress, nil
}

func (r *companyAddressRepository) Save(
	companyAddress entity.CompanyAddress,
) (companyAddressDB *entity.CompanyAddress, err error) {
	result := r.db.Save(&companyAddress)

	e := result.Error

	if e != nil {
		return nil, e
	}

	return &companyAddress, nil
}

func (r *companyAddressRepository) DeleteById(companyAddressId string) (str string, err error) {
	res := "Address successfully deleted"

	result := r.db.Delete(&entity.CompanyAddress{}, "company_address_id = ?", companyAddressId)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
