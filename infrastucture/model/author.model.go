package model

import "github.com/google/uuid"

type Author struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name" gorm:"type:varchar(100)"`
	Email string    `json:"email" gorm:"type:varchar(100)"`
	Books []Book    `json:"books" gorm:"foreig_key:BookID`
}
