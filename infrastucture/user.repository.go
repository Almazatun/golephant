package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	error_message "github.com/Almazatun/golephant/common/error-message"
	entity "github.com/Almazatun/golephant/infrastucture/entity"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	Save(user *entity.User) (registerUser *entity.User, err error)
	FindByEmail(email string) (user *entity.User, err error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *entity.User) (registerUser *entity.User, err error) {
	createUser := &entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	result := r.db.Create(&createUser)

	er := result.Error

	if er != nil {
		return nil, err
	}

	return createUser, nil
}

func (r *userRepository) FindByEmail(email string) (userDB *entity.User, err error) {
	var user entity.User

	result := r.db.First(&user, "email = ?", email)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.USER_NOT_FOUND)

		return nil, err
	}

	return &user, nil
}
