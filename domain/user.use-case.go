package usecase

import (
	"errors"
	"time"

	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
	"github.com/Almazatun/golephant/util"
	"github.com/dgrijalva/jwt-go"

	common "github.com/Almazatun/golephant/common"
	error_message "github.com/Almazatun/golephant/common/error-message"
	repository "github.com/Almazatun/golephant/infrastucture"
	"gopkg.in/go-playground/validator.v9"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	RegisterUser(registerUserInput *entity.User) (user *entity.User, err error)
	LogIn(logInInput input.LogIn) (resLogIn *ResLogIn, err error)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

type ResLogIn struct {
	Token             string
	ExperationTimeJWT time.Time
}

func (uc *userUseCase) RegisterUser(registerUserInput *entity.User) (user *entity.User, err error) {
	v := validator.New()
	e := v.Struct(registerUserInput)

	if e != nil {
		return nil, e
	}

	// Hashing user password
	hashedPassword, err := util.HashPassword(registerUserInput.Password)
	registerUserInput.Password = hashedPassword

	now := time.Now()
	registerUserInput.CreationTime = now
	registerUserInput.UpdateTime = now

	if err != nil {
		return nil, err
	}

	userDB, err := uc.userRepo.Save(registerUserInput)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}

func (uc *userUseCase) LogIn(logInInput input.LogIn) (resLogIn *ResLogIn, err error) {
	v := validator.New()
	e := v.Struct(logInInput)

	if e != nil {
		return nil, e
	}

	user, err := uc.userRepo.FindByEmail(logInInput.Email)

	if err != nil {
		return nil, err
	}

	isCorrectPassword := util.CheckPassword(logInInput.Password, user.Password)

	if !isCorrectPassword {
		newErr := errors.New(error_message.INCCORECT_PASSWORD)
		return nil, newErr
	}

	experationTimeJWT := time.Now().Add(time.Minute * 60)
	claims := common.Claims{
		UserEmail: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTimeJWT.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(common.JWT_KEY_BYTE)

	if err != nil {
		return nil, err
	}

	return &ResLogIn{Token: tokenString, ExperationTimeJWT: experationTimeJWT}, nil
}
