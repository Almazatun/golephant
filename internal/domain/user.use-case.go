package usecase

import (
	"errors"
	"time"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	"github.com/Almazatun/golephant/pkg/http/presentation/_type"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	"github.com/Almazatun/golephant/pkg/util"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	common "github.com/Almazatun/golephant/pkg/common"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	"gopkg.in/go-playground/validator.v9"
)

type userUseCase struct {
	userRepo               repository.UserRepo
	resetPasswordTokenRepo repository.ResetPasswordTokenRepo
}

type UserUseCase interface {
	Register(registerUserInput input.RegisterUserInput) (user *entity.User, err error)
	LogIn(logInInput input.LogInUserInput) (res *_type.ResLogIn[entity.User], err error)
	UpdateData(userId string, updateUserDataInput input.UpdateUserDataInput) (user *entity.User, err error)
	RequestResetPassword(userId string) (str *string, err error)
	ResetPassword(userId, resetPasswordToken, newPassword string) (str *string, err error)
}

func NewUserUseCase(
	userRepo repository.UserRepo,
	resetPasswordTokenRepo repository.ResetPasswordTokenRepo,
) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

type genericChan[T any] chan T

func (uc *userUseCase) Register(registerUserInput input.RegisterUserInput) (user *entity.User, err error) {
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

	generateJWT, err := jwt_gl.GenerateJWTStr(user.Email)

	if err != nil {
		return nil, err
	}

	return &_type.ResLogIn[entity.User]{
		Token:             generateJWT.Token,
		ExperationTimeJWT: generateJWT.ExperationTime,
		LogInEntityData:   *user}, nil
}

func (uc *userUseCase) UpdateData(userId string, updateUserDataInput input.UpdateUserDataInput) (user *entity.User, err error) {

	if isEmptyUpdateUserInput(updateUserDataInput) {
		return nil, nil
	}

	userDB, err := uc.userRepo.GetById(userId)

	if err != nil {
		return nil, err
	}

	updateUserData, err := updateUserColums(userDB, updateUserDataInput)

	if err != nil {
		return nil, err
	}

	res, err := uc.userRepo.Save(*updateUserData)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *userUseCase) RequestResetPassword(userId string) (str *string, err error) {
	user, err := uc.userRepo.GetById(userId)

	if err != nil {
		return nil, err
	}

	createResetPasswordTokenDB := entity.ResetPasswordToken{
		User:  *user,
		Token: util.GenerateRandomStr(),
	}

	resetPassworkTokenDB, err := uc.resetPasswordTokenRepo.Create(createResetPasswordTokenDB)

	if err != nil {
		return nil, err
	}

	common.SendEmail(user.Email, resetPassworkTokenDB.Token)

	resStr := "Pleace check you email"

	return &resStr, nil
}

func (uc *userUseCase) ResetPassword(userId, resetPasswordToken, newPassword string) (str *string, err error) {

	TokenDB, err := uc.resetPasswordTokenRepo.GetByToken(resetPasswordToken)

	if err != nil {
		return nil, err
	}

	if TokenDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	password := util.TrimWhiteSpace(newPassword)

	// Hashing user new password
	hashedPassword, err := util.HashPassword(password)

	if err != nil {
		return nil, err
	}

	TokenDB.User.Password = hashedPassword

	updateUser, err := uc.userRepo.Save(TokenDB.User)

	if err != nil {
		return nil, err
	}

	res := "Successfuly reset new password." + "User email: " + updateUser.Email

	return &res, nil
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
