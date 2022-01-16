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
	Save() (err error)
}

func InitUserRepo(db *gorm.DB) UserRepo {
	return &userRepository{db: db}
}

func (r *userRepository) Save() (err error) {
	fmt.Println("Save User")
	// query := "INSERT INTO users (email, username, password)  VALUES ($1, $2, $3)"

	return nil
}