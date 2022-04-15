package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Position struct {
	PositionID       uuid.UUID      `json:"position_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Description      string         `json:"name" gorm:"type:varchar(1000)"`
	Requirements     pq.StringArray `json:"requirements" gorm:"type:text[]"`
	Responsibilities pq.StringArray `json:"responsibilities" gorm:"type:text[]"`
	PositionType     string         `json:"position_type"  gorm:"type:text"`
	Salary           int            `json:"salary" gorm:"type:int"`
	CompanyID        uuid.UUID      `json:"company_id" gorm:"OnDelete:SET NULL; not NULL"`
	Company          Company
}
