package usecase

import (
	repository "github.com/Almazatun/golephant/infrastucture"
	"github.com/Almazatun/golephant/infrastucture/model"
	"gopkg.in/go-playground/validator.v9"
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
	v := validator.New()

	e := v.Struct(createUserInput)

	if e != nil {
		return nil, e
	}

	userDB, err := uc.userRepo.Save(createUserInput)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}
