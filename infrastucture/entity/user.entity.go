package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username     string    `json:"username" gorm:"type:varchar(100)"`
	Email        string    `json:"email" gorm:"type:varchar(100); not NULL; unique"`
	Password     string    `json:"password" gorm:"type:varchar(100); not NULL"`
	CreationTime time.Time `json:"creation_time" gorm:"type:date; not NULL"`
	UpdateTime   time.Time `json:"update_time" gorm:"type:date; not NULL"`
	Mobile       string    `json:"mobile" gorm:"type:varchar(30); default:NULL"`
	Status       string    `json:"status" gorm:"type:text; default:0"`
	Resume       []Resume  `gorm:"ForeignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
