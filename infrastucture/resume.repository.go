package repository

import (
	"errors"

	error_message "github.com/Almazatun/golephant/common/error-message"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/jinzhu/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

type ResumeRepo interface {
	Create(resume entity.Resume) (resumeDB *entity.Resume, err error)
	DeleteById(resumeId string) (str string, err error)
	FindById(resumeId string) (resumeDB *entity.Resume, err error)
	Update(resume entity.Resume) (resumeDB *entity.Resume, err error)
}

func NewResumeRepo(db *gorm.DB) ResumeRepo {
	return &resumeRepository{
		db: db,
	}
}

func (r *resumeRepository) Create(resume entity.Resume) (resumeDB *entity.Resume, err error) {
	var createResume entity.Resume

	result := r.db.Create(&resume)

	er := result.Error

	if er != nil {
		return nil, err
	}

	createResume = resume

	return &createResume, nil
}

func (r *resumeRepository) DeleteById(resumeId string) (str string, err error) {
	res := "Resume successfully deleted"

	result := r.db.Delete(&entity.Resume{}, "resume_id = ?", resumeId)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}

func (r *resumeRepository) FindById(resumeId string) (resumeDB *entity.Resume, err error) {
	var resume entity.Resume

	result := r.db.Preload("UserEducations").Preload("UserExperiences").First(&resume, "resume_id = ?", resumeId)

	r.db.Model(resume).Related(&resume.User)

	er := result.Error

	if er != nil {
		err := errors.New(error_message.RESUME_NOT_FOUND)
		return nil, err
	}

	return &resume, nil
}

func (r *resumeRepository) Update(updateResume entity.Resume) (resumeDB *entity.Resume, err error) {
	var result entity.Resume

	res := r.db.Model(&result).Preload("UserEducations").Preload("UserExperiences").Updates(updateResume)

	e := res.Error

	if e != nil {
		return nil, e
	}

	r.db.Model(result).Related(&result.User)

	return &result, nil
}
