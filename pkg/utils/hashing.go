package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a hash for the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPassword []byte, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
	return err == nil
}
