package core

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost = 10
)

func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func IsPasswordValid(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
