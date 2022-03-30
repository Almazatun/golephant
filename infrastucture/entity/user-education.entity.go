package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserEducation struct {
	UserEducationID uuid.UUID `json:"user_education_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	CreationTime    time.Time `json:"creation_time" gorm:"type:date; not NULL"`
	UpdateTime      time.Time `json:"update_time" gorm:"type:date; not NULL"`
	DegreePlacement string    `json:"degree_placement" gorm:"type:varchar(100);default:NULL"`
	City            string    `json:"city" gorm:"type:varchar(100);not NULL"`
	ResumeID        uuid.UUID `json:"resume_id" gorm:"not NULL"`
	Resume          Resume
}
