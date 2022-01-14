package model

import "github.com/google/uuid"

type Book struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title  string    `json:"title" gorm:"type:varchar(100)"`
	Author Author    `json:"author" gorm: "foreing_key: AuthorID`
}
