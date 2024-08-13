package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
