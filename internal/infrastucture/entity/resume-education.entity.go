package entity

import (
	"time"

	"github.com/google/uuid"
)

type ResumeEducation struct {
	ResumeEducationID uuid.UUID `json:"resume_education_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	DegreePlacement   string    `json:"degree_placement" gorm:"type:varchar(100);default:0;"`
	City              string    `json:"city" gorm:"type:varchar(100);not NULL"`
	ResumeID          uuid.UUID `json:"resume_id" gorm:"not NULL"`
}
