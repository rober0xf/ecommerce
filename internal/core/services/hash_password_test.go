package services

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash, got nothing")
	}

	if hash == password {
		t.Errorf("expected hash different from password")
	}

	t.Run("comparing password | valid match", func(t *testing.T) {
		if err := CheckPassword(hash, password); err != nil {
			t.Errorf("expected password to match, but got: %v", err)
		}
	})

	t.Run("comparing password | invalid match", func(t *testing.T) {
		if err := CheckPassword(hash, "wrongpassword"); err == nil {
			t.Errorf("expected password to not match")
		}
	})
}
