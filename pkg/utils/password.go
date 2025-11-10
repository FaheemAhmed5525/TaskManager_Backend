package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain text password and returns the bcrypt hashed password
func HashPassword(password string) (string, error) {
	hashed, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		return "", error
	}
	return string(hashed), nil
}

// CheckPassword compares a plain text password with a bcrypt hashed password
func CheckHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
