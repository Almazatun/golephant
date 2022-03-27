package repository

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	error_message "github.com/Almazatun/golephant/common/error-message"
	entity "github.com/Almazatun/golephant/infrastucture/entity"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	Create(user entity.User) (registerUser *entity.User, err error)
	FindByEmail(email string) (user *entity.User, err error)
	FindById(userId string) (user *entity.User, err error)
	Update(updateUserData entity.User) (updateUser *entity.User, err error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (registerUser *entity.User, err error) {
	var createUser entity.User

	result := r.db.Create(&user)

	er := result.Error

	if er != nil {
		return nil, err
	}

	createUser = user

	return &createUser, nil
}

func (r *userRepository) FindByEmail(email string) (userDB *entity.User, err error) {
	var user entity.User

	result := r.db.First(&user, "email = ?", email)

	fmt.Println(user)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_NOT_FOUND)

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindById(userId string) (userDB *entity.User, err error) {
	var user entity.User

	result := r.db.First(&user, "user_id = ?", userId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_NOT_FOUND)

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(updateUserData entity.User) (updateUser *entity.User, err error) {
	var user entity.User

	result := r.db.Model(&user).Updates(updateUserData)

	e := result.Error

	if e != nil {
		return nil, e
	}

	return &user, nil

}
