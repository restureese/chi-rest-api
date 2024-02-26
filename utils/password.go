package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func VerifyPassword(pwdHash string, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
	return err
}
