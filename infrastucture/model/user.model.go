package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username" gorm:"type:varchar(100)"`
	Email    string    `json:"email" validate:"required,email,omitempty" gorm:"type:varchar(100); not NULL; unique"`
	Password string    `json:"password" validate:"required,max=20,min=7" gorm:"type:varchar(100); not NULL"`
}
