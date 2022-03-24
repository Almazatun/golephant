package entity

import (
	"github.com/google/uuid"
)

type CompanyAddress struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4(); primary_key"`
	Address       string    `json:"address" validate:"required,address" gorm:"type: varchar(200); not NULL;"`
	IsBaseAddress bool      `json:"is_base_adress" validate:"required" gorm:"type: boolean"`
	CompanyID     uuid.UUID `json:"company_id" gorm:"OnDelete:SET NULL; not NULL"`
	Company       Company
}
