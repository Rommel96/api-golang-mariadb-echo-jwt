package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(passwordEncrypted, password string) error {
	if password == "" {
		return bcrypt.ErrMismatchedHashAndPassword
	}
	return bcrypt.CompareHashAndPassword([]byte(passwordEncrypted), []byte(password))
}
