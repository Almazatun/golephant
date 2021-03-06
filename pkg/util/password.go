package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (str string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) (res bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
