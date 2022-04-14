package usecase

import (
	"errors"
	"time"

	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/_type"
	"github.com/Almazatun/golephant/presentation/input"
	"github.com/Almazatun/golephant/util"

	common "github.com/Almazatun/golephant/common"
	error_message "github.com/Almazatun/golephant/common/error-message"
	repository "github.com/Almazatun/golephant/infrastucture"
	"gopkg.in/go-playground/validator.v9"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	RegisterUser(registerUserInput input.RegisterUserInput) (user *entity.User, err error)
	LogIn(logInInput input.LogInUserInput) (res *_type.ResLogIn[entity.User], err error)
	UpdateUserData(userId string, updateUserDataInput input.UpdateUserDataInput) (user *entity.User, err error)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) RegisterUser(registerUserInput input.RegisterUserInput) (user *entity.User, err error) {
	// Validate register user input
	v := validator.New()
	e := v.Struct(registerUserInput)

	if e != nil {
		return nil, e
	}

	registerUser := registerUserColums(registerUserInput)

	// Hashing user password
	hashedPassword, err := util.HashPassword(registerUserInput.Password)

	if err != nil {
		return nil, err
	}

	registerUser.Password = hashedPassword

	now := time.Now()
	registerUser.CreationTime = now
	registerUser.UpdateTime = now

	userDB, err := uc.userRepo.Create(registerUser)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}

func (uc *userUseCase) LogIn(logInInput input.LogInUserInput) (res *_type.ResLogIn[entity.User], err error) {
	// Validate register user input
	v := validator.New()
	e := v.Struct(logInInput)

	if e != nil {
		return nil, e
	}

	user, err := uc.userRepo.GetByEmail(logInInput.Email)

	if err != nil {
		return nil, err
	}

	isCorrectPassword := util.CheckPassword(logInInput.Password, user.Password)

	if !isCorrectPassword {
		newErr := errors.New(error_message.INCCORECT_PASSWORD)
		return nil, newErr
	}

	generateJWT, err := common.GenerateJWTStr(user.Email)

	if err != nil {
		return nil, err
	}

	return &_type.ResLogIn[entity.User]{
		Token:             generateJWT.Token,
		ExperationTimeJWT: generateJWT.ExperationTime,
		LogInEntityData:   *user}, nil
}

func (uc *userUseCase) UpdateUserData(userId string, updateUserDataInput input.UpdateUserDataInput) (user *entity.User, err error) {

	if isEmptyUpdateUserInput(updateUserDataInput) {
		return nil, nil
	}

	userDB, err := uc.userRepo.GetByEmail(userId)

	if err != nil {
		return nil, err
	}

	updateUserData, err := updateUserColums(userDB, updateUserDataInput)

	if err != nil {
		return nil, err
	}

	res, err := uc.userRepo.Update(*updateUserData)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func registerUserColums(registerUserInput input.RegisterUserInput) (registerUser entity.User) {

	if registerUserInput.Email != "" {
		registerUser.Email = registerUserInput.Email
	}

	if registerUserInput.Mobile != "" {
		registerUser.Mobile = registerUserInput.Mobile
	}

	if registerUserInput.Status != "" {
		registerUser.Status = registerUserInput.Status
	}

	if registerUserInput.Username != "" {
		registerUser.Username = registerUserInput.Username
	}

	return registerUser
}

func updateUserColums(userDB *entity.User, updateUserDataInput input.UpdateUserDataInput) (updateUserData *entity.User, err error) {

	if updateUserDataInput.Email != "" {
		userDB.Email = updateUserDataInput.Email
	}

	if updateUserDataInput.Mobile != "" {
		userDB.Mobile = updateUserDataInput.Mobile
	}

	if updateUserDataInput.Username != "" {
		userDB.Username = updateUserDataInput.Username
	}

	if updateUserDataInput.Password != "" {
		// Hashing user password
		hashedPassword, err := util.HashPassword(updateUserDataInput.Password)

		if err != nil {
			return nil, err
		}

		userDB.Password = hashedPassword
	}

	return userDB, nil
}

func isEmptyUpdateUserInput(updateUserDataInput input.UpdateUserDataInput) bool {
	return (input.UpdateUserDataInput{}) == updateUserDataInput
}
