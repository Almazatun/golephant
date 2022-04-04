package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserExperience struct {
	UserExperienceID uuid.UUID `json:"user_experience_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	CompanyName      string    `json:"company_name" gorm:"type:varchar(100); not NULL"`
	Position         string    `json:"position" gorm:"type:varchar(100); default:null"`
	City             string    `json:"city" gorm:"type:varchar(100); not NULL"`
	ResumeID         uuid.UUID `json:"resume_id" gorm:"not NULL"`
	Resume           Resume
}
