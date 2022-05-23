package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Resume struct {
	ResumeID        uuid.UUID      `json:"resume_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName       string         `json:"first_name" gorm:"type:varchar(100); not NULL"`
	LastName        string         `json:"last_name" gorm:"type:varchar(100); not NULL"`
	DateOfBirght    time.Time      `json:"date_of_birght" gorm:"type:date; not NULL"`
	Gender          string         `json:"gender" gorm:"type:text; default:0"`
	Citizenship     string         `json:"citizenship" gorm:"type:text; default:0;"`
	DesiredPosition string         `json:"desired_position" gorm:"type:varchar(100);index:idx_member"`
	SubwayStation   string         `json:"subway_station" gorm:"type:text; default:0;"`
	Specialization  string         `json:"specialization" gorm:"type:text; default:0;"`
	WorkMode        string         `json:"work_mode" gorm:"type:text;default:0; not NULL"`
	About           string         `json:"about" gorm:"type:text"`
	Tags            pq.StringArray `json:"tags" gorm:"type:text[]"`
	Status          string         `json:"status" gorm:"type:text"`
	CreationTime    time.Time      `json:"creation_time" gorm:"type:date; not NULL"`
	UpdateTime      time.Time      `json:"update_time" gorm:"type:date; not NULL"`
	UserID          uuid.UUID      `json:"user_id" gorm:"not NULL; index:idx_member"`
	User            User
	UserExperience []UserExperience `gorm:"ForeignKey:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserEducation  []UserEducation  `gorm:"ForeignKey:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
