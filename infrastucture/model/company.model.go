package model

import "github.com/google/uuid"

type Company struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string     `json:"title" validate:"required" gorm:"type:varchar(100); not NULL; unique"`
	Phone     string     `json:"phone" validate:"required" gorm:"type:varchar(100); not NULL"`
	Address   string     `json:"address" gorm:"type:varchar(100)"`
	Positions []Position `gorm:"foreignKey:PositionID"`
}
