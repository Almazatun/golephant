package entity

import (
	"time"

	"github.com/google/uuid"
)

type ResetPasswordToken struct {
	ResetPasswordTokenID uuid.UUID  `json:"reset_password_token_id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Token                string     `json:"token" gorm:"type:varchar(32)"`
	CompanyID            *uuid.UUID `json:"company_id"`
	Company              Company
	UserID               *uuid.UUID `json:"user_id"`
	User                 User
	CreationTime         time.Time `json:"creation_time" gorm:"type:date; not NULL"`
}
