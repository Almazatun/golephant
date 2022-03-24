package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Position struct {
	Id               uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Description      string         `json:"name" gorm:"type:varchar(1000)"`
	Requirements     pq.StringArray `json:"requirements" gorm:"type:text[]" validate:"required"`
	Responsibilities pq.StringArray `json:"responsibilities" gorm:"type:text[]" validate:"required"`
	PositionTypeID   *uuid.UUID
	PositionType     PositionType
}
