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
	// CreateUser(ctx context.Context, user *model.User) (err error)
	Save(user *model.User) (createdUser *model.User, err error)
	FindByEmail(email string) (err error)
}

func InitUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *model.User) (createdUser *model.User, err error) {
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

func (r *userRepository) FindByEmail(email string) (err error) {
	user := r.db.First(&model.User{}, "email = ?", email)

	fmt.Println(user)

	return nil
}
