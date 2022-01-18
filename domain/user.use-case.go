package usecase

import (
	repository "github.com/Almazatun/golephant/infrastucture"
	"github.com/Almazatun/golephant/infrastucture/model"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	CreateUser(createUserInput *model.User) (user *model.User, err error)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser(createUserInput *model.User) (user *model.User, err error) {
	userDB, err := uc.userRepo.Save(createUserInput)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}
