package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Position struct {
	PositionID       uuid.UUID      `json:"position_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Description      string         `json:"description" gorm:"type:varchar(1000)"`
	Requirements     pq.StringArray `json:"requirements" gorm:"type:text[]"`
	Responsibilities pq.StringArray `json:"responsibilities" gorm:"type:text[]"`
	PositionType     string         `json:"position_type"  gorm:"type:text"`
	Salary           *int           `json:"salary" gorm:"type:int"`
	CompanyID        uuid.UUID      `json:"company_id" gorm:"not NULL"`
	Status           string         `json:"status" gorm:"text; not NULL;default:OPEN;"`
	CreationTime     time.Time      `json:"creation_time" gorm:"type:date; not NULL;"`
	UpdateTime       time.Time      `json:"update_time" gorm:"type:date; not NULL;"`
}
