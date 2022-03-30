package repository

import (
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/jinzhu/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

type ResumeRepo interface {
	Create(resume entity.Resume) (res *entity.Resume, err error)
}

func NewResumeRepo(db *gorm.DB) ResumeRepo {
	return &resumeRepository{
		db: db,
	}
}

func (r *resumeRepository) Create(resume entity.Resume) (res *entity.Resume, err error) {
	var createResume entity.Resume

	result := r.db.Create(&resume)

	er := result.Error

	if er != nil {
		return nil, err
	}

	createResume = resume

	return &createResume, nil
}
