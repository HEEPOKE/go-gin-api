package common

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
