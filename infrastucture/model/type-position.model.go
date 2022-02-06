package model

import "github.com/google/uuid"

type TypePosition struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title string    `json:"title" validate:"required" gorm:"type:varchar(100); not null; unique"`
}
