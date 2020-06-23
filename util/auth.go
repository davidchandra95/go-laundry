package util

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := Hash(password)
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return string(hashedPassword), nil
}
