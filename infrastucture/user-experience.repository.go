package repository

import (
	"errors"

	error_message "github.com/Almazatun/golephant/common/error-message"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/jinzhu/gorm"
)

type userExperienceRepository struct {
	db *gorm.DB
}

type UserExperienceRepo interface {
	FindById(userExperienceId string) (userExperienceDB *entity.UserExperience, err error)
	DeleteById(userExperienceId string) (str string, err error)
}

func NewUserExperienceRepo(db *gorm.DB) UserExperienceRepo {
	return &userExperienceRepository{
		db: db,
	}
}

func (r *userExperienceRepository) FindById(userExperienceId string) (userExperienceDB *entity.UserExperience, err error) {
	var userEducation entity.UserExperience

	result := r.db.First(&userEducation, "user_experience_id = ?", userExperienceId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_EXPERIENCE_NOT_FOUND)

		return nil, err
	}

	return &userEducation, nil
}

func (r *userExperienceRepository) DeleteById(userExperienceId string) (str string, err error) {
	res := "User experience successfully deleted in resume"

	result := r.db.Delete(&entity.UserExperience{}, "user_experience_id = ?", userExperienceId)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
