package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Resume struct {
	ResumeID       uuid.UUID      `json:"resume_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title          string         `json:"title" gorm:"type:varchar(100); not NULL; index:idx_member"`
	Specialization string         `json:"specialization" gorm:"type:text;default:0"`
	WorkMode       string         `json:"work_mode" gorm:"type:text;default:0; not NULL"`
	About          string         `json:"about" gorm:"type:text;default:0;"`
	Tags           pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreationTime   time.Time      `json:"creation_time" gorm:"type:date; not NULL"`
	UpdateTime     time.Time      `json:"update_time" gorm:"type:date; not NULL"`
	UserID         uuid.UUID      `json:"user_id" gorm:"not NULL; index:idx_member"`
	User           User
	UserExperience []UserExperience `gorm:"ForeignKey:ResumeID;references:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserEducation  []UserEducation  `gorm:"ForeignKey:ResumeID;references:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
