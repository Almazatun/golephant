package repository

import (
	"errors"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type resumeEducationRepository struct {
	db *gorm.DB
}

type ResumeEducationRepo interface {
	GetById(id string) (resumeEducationDB *entity.ResumeEducation, err error)
	DeleteById(id string) (str string, err error)
}

func NewResumeEducationRepo(db *gorm.DB) ResumeEducationRepo {
	return &resumeEducationRepository{
		db: db,
	}
}

func (r *resumeEducationRepository) GetById(id string) (resumeEducationDB *entity.ResumeEducation, err error) {
	var education entity.ResumeEducation

	result := r.db.First(&education, "resume_education_id = ?", id)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_EDUCATION_NOT_FOUND)

		return nil, err
	}

	return &education, nil
}

func (r *resumeEducationRepository) DeleteById(id string) (str string, err error) {
	res := "Education successfully deleted in resume"

	result := r.db.Delete(&entity.ResumeEducation{}, "resume_education_id = ?", id)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
