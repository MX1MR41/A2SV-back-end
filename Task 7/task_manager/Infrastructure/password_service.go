package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

// Functionality that checks if hashing the plainPassword gives the hashedPassword
func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
