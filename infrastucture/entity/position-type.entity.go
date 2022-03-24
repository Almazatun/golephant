package entity

import "github.com/google/uuid"

type PositionType struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title string    `json:"title" validate:"required" gorm:"type:varchar(100); not NULL; unique"`
}
