package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	model "github.com/Almazatun/golephant/infrastucture/model"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	Save(user *model.User) (registerUser *model.User, err error)
	FindByEmail(email string) (user *model.User, err error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *model.User) (registerUser *model.User, err error) {
	createUser := &model.User{
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

func (r *userRepository) FindByEmail(email string) (userDB *model.User, err error) {
	user := r.db.First(&model.User{})
	fmt.Println(user)

	return userDB, nil
}
