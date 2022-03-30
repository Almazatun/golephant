package repository

import (
	"errors"

	error_message "github.com/Almazatun/golephant/common/error-message"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/jinzhu/gorm"
)

type userEducationRepository struct {
	db *gorm.DB
}

type UserEducationRepo interface {
	FindById(userEducationId string) (userEducationDB *entity.UserEducation, err error)
	DeleteById(userEducationId string) (str string, err error)
}

func NewUserEducationRepo(db *gorm.DB) UserEducationRepo {
	return &userEducationRepository{
		db: db,
	}
}

func (r *userEducationRepository) FindById(userEducationId string) (userEducationDB *entity.UserEducation, err error) {
	var userEducation entity.UserEducation

	result := r.db.First(&userEducation, "user_education_id = ?", userEducationId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_EDUCATION_NOT_FOUND)

		return nil, err
	}

	return &userEducation, nil
}

func (r *userEducationRepository) DeleteById(userEducationId string) (str string, err error) {
	res := "User education successfully deleted in resume"

	result := r.db.Delete(&entity.UserEducation{}, "user_education_id = ?", userEducationId)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
