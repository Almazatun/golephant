package repository

import (
	"errors"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/jinzhu/gorm"
)

type resetPasswordTokenRepo struct {
	db *gorm.DB
}

type ResetPasswordTokenRepo interface {
	Create(resetPasswordToken entity.ResetPasswordToken) (resetPasswordTokenDB *entity.ResetPasswordToken, err error)
	Delete(resetPasswordTokenId string) (str *string, err error)
	GetByToken(token string) (resetPasswordTokenDB *entity.ResetPasswordToken, err error)
}

func NewResetPasswordTokenRepo(db *gorm.DB) ResetPasswordTokenRepo {
	return &resetPasswordTokenRepo{
		db: db,
	}
}

func (r *resetPasswordTokenRepo) Create(
	resetPasswordToken entity.ResetPasswordToken,
) (resetPasswordTokenDB *entity.ResetPasswordToken, err error) {
	result := r.db.Create(&resetPasswordToken)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return &resetPasswordToken, nil
}

func (r *resetPasswordTokenRepo) Delete(resetPasswordTokenId string) (str *string, err error) {
	res := "Reset Password token successfully deleted"

	result := r.db.Delete(&entity.ResetPasswordToken{}, "reset_password_token_id = ?", resetPasswordTokenId)

	er := result.Error

	if er != nil {
		return nil, er
	}

	return &res, nil
}

func (r *resetPasswordTokenRepo) GetByToken(token string) (resetPasswordTokenDB *entity.ResetPasswordToken, err error) {
	var resetPasswordToken entity.ResetPasswordToken

	result := r.db.First(&resetPasswordToken, "token = ?", token)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.INVALID_RESET_PASSWORD_TOKEN)

		return nil, err
	}

	return &resetPasswordToken, nil
}
