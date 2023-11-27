package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword return hashed value of password
func HashedPassword(passowrd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
