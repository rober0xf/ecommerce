package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash_pass), nil
}
