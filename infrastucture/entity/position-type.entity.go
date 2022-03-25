package entity

import "github.com/google/uuid"

type PositionType struct {
	PositionTypeID uuid.UUID `json:"position_type_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title          string    `json:"title" validate:"required" gorm:"type:varchar(100); not NULL; unique"`
}
