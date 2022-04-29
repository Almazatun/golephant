package entity

import (
	"github.com/google/uuid"
)

type CompanyAddress struct {
	CompanyAddressID uuid.UUID `json:"company_address_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title            string    `json:"title" gorm:"type: varchar(200); not NULL"`
	IsBaseAddress    bool      `json:"is_base_adress" gorm:"type: boolean"`
	CompanyID        uuid.UUID `json:"company_id" gorm:"not NULL"`
}
