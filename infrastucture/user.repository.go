package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	// CreateUser(ctx context.Context, user *model.User) (err error)
	CreateUser() (err error)
}

func InitUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser() (err error) {
	fmt.Println("CreateUser")
	// query := "INSERT INTO users (email, username, password)  VALUES ($1, $2, $3)"

	return nil
}
