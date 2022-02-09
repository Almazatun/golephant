package usecase

import (
	"github.com/Almazatun/golephant/presentation/input"
	util "github.com/Almazatun/golephant/util"

	repository "github.com/Almazatun/golephant/infrastucture"
	"github.com/Almazatun/golephant/infrastucture/model"
	"gopkg.in/go-playground/validator.v9"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	RegisterUser(registerUserInput *model.User) (user *model.User, err error)
	LogIn(logInInput *input.LogIn) (str string, err error)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) RegisterUser(registerUserInput *model.User) (user *model.User, err error) {
	v := validator.New()
	e := v.Struct(registerUserInput)

	if e != nil {
		return nil, e
	}

	hashedPassword, err := util.HashPassword(registerUserInput.Password)
	registerUserInput.Password = hashedPassword

	if err != nil {
		return nil, err
	}

	userDB, err := uc.userRepo.Save(registerUserInput)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}

func (uc *userUseCase) LogIn(logInInput *input.LogIn) (str string, err error) {
	v := validator.New()
	e := v.Struct(logInInput)

	if e != nil {
		return "", e
	}

	uc.userRepo.FindByEmail(logInInput.Email)

	return "", nil
}
