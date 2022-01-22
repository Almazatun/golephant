package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username" gorm:"type:varchar(100)"`
	Email    string    `json:"email" validate:"require,email" gorm:"type:varchar(100); not null; unique"`
	Password string    `json:"password" validate:"require,email" gorm:"type:varchar(100); not null"`
}
