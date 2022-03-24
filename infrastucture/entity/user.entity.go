package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID        `json:"id" gorm:"type:uuid;default:uuid_generate_v4(); primary_key"`
	Username     string           `json:"username" gorm:"type:varchar(100)"`
	Email        string           `json:"email" validate:"required,email,omitempty" gorm:"type:varchar(100); not NULL; unique"`
	Password     string           `json:"password" validate:"required,max=20,min=7" gorm:"type:varchar(100); not NULL"`
	CreationTime time.Time        `json:"creation_time" gorm:"type:date; not NULL"`
	UpdateTime   time.Time        `json:"update_time" gorm:"type:date; not NULL"`
	Age          string           `json:"age" gorm:"type:text; default:0"`
	Mobile       string           `json:"mobile" gorm:"type:varchar(30); default:NULL"`
	Status       string           `json:"status" gorm:"type:text; default:0"`
	City         string           `json:"city" gorm:"type:varchar(100);default:NUll"`
	Experiences  []UserExperience `json:"experiences" gorm:"ForeignKey:UserID;references:UserID;"`
	Educations   []UserEducation  `json:"educations" gorm:"ForeignKey:UserID;references:UserID;"`
}
