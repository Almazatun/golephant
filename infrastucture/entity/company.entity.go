package entity

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	CompanyID      uuid.UUID        `json:"company_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title          string           `json:"title" validate:"required" gorm:"type:varchar(100); not NULL; unique"`
	Email          string           `json:"email" validate:"required,email,omitempty" gorm:"type:varchar(100); not NULL;"`
	Phone          string           `json:"phone" validate:"required" gorm:"type:varchar(100); not NULL;"`
	CreationTime   time.Time        `json:"creation_time" gorm:"type:date; not NULL;"`
	UpdateTime     time.Time        `json:"update_time" gorm:"type:date; not NULL;"`
	CompanyAddress []CompanyAddress `gorm:"ForeignKey:CompanyID;references:CompanyID;"`
}
