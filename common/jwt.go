package common

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")
var JWT_KEY_BYTE = []byte(secretKey)

type Claims struct {
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}
