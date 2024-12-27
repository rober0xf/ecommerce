package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash_pass), nil
}

func CheckPassword(hashed_password, password string) bool {
	// returns nil on success or error on failure
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err == nil {
		return true
	}

	return false
}
