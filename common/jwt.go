package common

import (
	"errors"
	"os"

	error_message "github.com/Almazatun/golephant/common/error-message"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")
var SET_COOKIE_PATH = os.Getenv("SET_COOKIE_PATH")
var JWT_KEY_BYTE = []byte(secretKey)

const (
	HTTP_COOKIE = "Token"
)

type Claims struct {
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}

func IsValidJWTStr(tokenStr string) (res bool, err error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY_BYTE, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			newErr := errors.New(error_message.UNAUTHORIZED)

			return false, newErr
		}
		errMes := "Bad request"
		newErr := errors.New(errMes)

		return false, newErr
	}

	if !token.Valid {
		newErr := errors.New(error_message.UNAUTHORIZED)

		return false, newErr
	}

	return true, nil
}
