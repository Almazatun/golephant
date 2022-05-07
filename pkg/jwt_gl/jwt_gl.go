package jwt_gl

import (
	"errors"
	"os"
	"time"

	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")
var SET_COOKIE_PATH = os.Getenv("SET_COOKIE_PATH")
var JWT_KEY_BYTE = []byte(secretKey)

const (
	HTTP_COOKIE = "Token"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWT struct {
	Token          string
	ExperationTime time.Time
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

func GenerateJWTStr(email string) (res *JWT, err error) {
	experationTimeJWT := time.Now().Add(time.Minute * 60)
	claims := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTimeJWT.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY_BYTE)

	if err != nil {
		return nil, err
	}

	return &JWT{Token: tokenString, ExperationTime: experationTimeJWT}, nil

}
