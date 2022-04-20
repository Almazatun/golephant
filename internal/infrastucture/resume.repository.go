package repository

import (
	"errors"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

type ResumeRepo interface {
	ListByUserId(userId string) (resumeDB *[]entity.Resume, err error)
	Create(resume entity.Resume) (resumeDB *entity.Resume, err error)
	DeleteById(resumeId string) (str string, err error)
	GetById(resumeId string) (resumeDB *entity.Resume, err error)
	Save(resume entity.Resume) (resumeDB *entity.Resume, err error)
}

func NewResumeRepo(db *gorm.DB) ResumeRepo {
	return &resumeRepository{
		db: db,
	}
}

func (r *resumeRepository) ListByUserId(userId string) (resumeDB *[]entity.Resume, err error) {
	var list []entity.Resume

	result := r.db.
		Preload("UserEducations").
		Preload("UserExperiences").
		// Preload("User").
		Find(&list, "user_id = ?", userId)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return &list, nil
}

func (r *resumeRepository) Create(resume entity.Resume) (resumeDB *entity.Resume, err error) {

	result := r.db.Create(&resume)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return &resume, nil
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

func (r *resumeRepository) GetById(resumeId string) (resumeDB *entity.Resume, err error) {
	var resume entity.Resume

	result := r.db.
		Preload("UserEducations").
		Preload("UserExperiences").
		// Preload("User").
		First(&resume, "resume_id = ?", resumeId)

	er := result.Error

	if er != nil {
		err := errors.New(error_message.RESUME_NOT_FOUND)
		return nil, err
	}

	return &resume, nil
}

func (r *resumeRepository) Save(resume entity.Resume) (resumeDB *entity.Resume, err error) {
	var result entity.Resume

	res := r.db.Save(&result).
		Preload("UserEducations").
		Preload("UserExperiences")

	e := res.Error

	if e != nil {
		return nil, e
	}

	return &result, nil
}
