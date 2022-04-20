package repository

import (
	"errors"

	"gorm.io/gorm"

	entity "github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	Create(user entity.User) (userDB *entity.User, err error)
	GetByEmail(email string) (userDB *entity.User, err error)
	GetById(userId string) (userDB *entity.User, err error)
	Save(user entity.User) (userDB *entity.User, err error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (registerUser *entity.User, err error) {
	result := r.db.Create(&user)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (userDB *entity.User, err error) {
	var user entity.User

	result := r.db.First(&user, "email = ?", email)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_NOT_FOUND)

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetById(userId string) (userDB *entity.User, err error) {
	var user entity.User

	result := r.db.First(&user, "user_id = ?", userId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_NOT_FOUND)

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Save(user entity.User) (updateUser *entity.User, err error) {

	result := r.db.Save(&user)

	e := result.Error

	if e != nil {
		return nil, e
	}

	return &user, nil

}
