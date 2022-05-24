package repository

import (
	"errors"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type resumeExperienceRepository struct {
	db *gorm.DB
}

type ResumeExperienceRepo interface {
	GetById(id string) (resumeExperienceDB *entity.ResumeExperience, err error)
	DeleteById(id string) (str string, err error)
}

func NewResumeExperienceRepo(db *gorm.DB) ResumeExperienceRepo {
	return &resumeExperienceRepository{
		db: db,
	}
}

func (r *resumeExperienceRepository) GetById(id string) (resumeExperienceDB *entity.ResumeExperience, err error) {
	var education entity.ResumeExperience

	result := r.db.First(&education, "resume_experience_id = ?", id)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_EXPERIENCE_NOT_FOUND)

		return nil, err
	}

	return &education, nil
}

func (r *resumeExperienceRepository) DeleteById(id string) (str string, err error) {
	res := "Experience successfully deleted in resume"

	result := r.db.Delete(&entity.ResumeExperience{}, "resume_experience_id = ?", id)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
