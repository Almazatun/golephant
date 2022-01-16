package usecase

import (
	"fmt"

	repository "github.com/Almazatun/golephant/infrastucture"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	CreateUser()
}

func InitUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser() {
	fmt.Println("CreateUser UseCase")
	uc.userRepo.Save()
}
