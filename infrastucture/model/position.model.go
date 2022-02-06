package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Position struct {
	Id               uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Description      string         `json:"name" gorm:"type:varchar(1000)"`
	Requirements     pq.StringArray `gorm:"type:text[]" validate:"required"`
	Responsibilities pq.StringArray `gorm:"type:text[]" validate:"required"`
	Type             TypePosition   `gorm:"foreignKey:TypePositionID"`
}
