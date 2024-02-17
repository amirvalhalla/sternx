package util

import "golang.org/x/crypto/scrypt"

func EncryptPassword(password string, salt []byte) ([]byte, error) {
	hashedPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	return hashedPassword, err
}
