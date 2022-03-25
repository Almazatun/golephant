package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
	"github.com/Almazatun/golephant/util"
	"github.com/dgrijalva/jwt-go"

	repository "github.com/Almazatun/golephant/infrastucture"
	"gopkg.in/go-playground/validator.v9"
)

type userUseCase struct {
	userRepo repository.UserRepo
}

type UserUseCase interface {
	RegisterUser(registerUserInput *entity.User) (user *entity.User, err error)
	LogIn(logInInput input.LogIn) (str string, err error)
}

var secretKey = os.Getenv("JWT_SECRET_KEY")
var jwtKey = []byte(secretKey)

type Claims struct {
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}

func NewUserUseCase(userRepo repository.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
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

func (uc *userUseCase) LogIn(logInInput input.LogIn) (str string, err error) {
	v := validator.New()
	e := v.Struct(logInInput)

	if e != nil {
		return "", e
	}

	user, err := uc.userRepo.FindByEmail(logInInput.Email)

	if err != nil {
		return "", err
	}

	isCorrectPassword := util.CheckPassword(logInInput.Password, user.Password)

	if !isCorrectPassword {
		newErr := errors.New("Incorrect password")
		return "", newErr
	}

	experationTimeJWT := time.Now().Add(time.Minute * 60)
	claims := Claims{
		UserEmail: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTimeJWT.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
