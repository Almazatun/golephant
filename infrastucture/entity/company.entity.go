package entity

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	CompanyID      uuid.UUID        `json:"company_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title          string           `json:"title" gorm:"type:varchar(100); not NULL; unique"`
	Email          string           `json:"email" gorm:"type:varchar(100); not NULL; unique"`
	Password       string           `json:"password" gorm:"type:varchar(100); not NULL"`
	Phone          string           `json:"phone" gorm:"type:varchar(100);"`
	CreationTime   time.Time        `json:"creation_time" gorm:"type:date; not NULL;"`
	UpdateTime     time.Time        `json:"update_time" gorm:"type:date; not NULL;"`
	CompanyAddress []CompanyAddress `gorm:"ForeignKey:CompanyID;references:CompanyID;"`
}
