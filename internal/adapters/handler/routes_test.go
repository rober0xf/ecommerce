package handler

import (
	"bytes"
	"ecommerce/internal/core/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user *models.User) error {
	return nil
}

func TestUserHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore) // we need to define all the methods of the interface

	t.Run("should fail if the payload is invalid", func(t *testing.T) {
		payload := models.Payload{
			Name:     "namee",
			LastName: "lastname",
			Email:    "invalid_email",
			Password: "123434324",
		}

		marshalled, err := json.Marshal(payload) // convert the struct to a valid json
		if err != nil {
			t.Errorf("failed to marshal payload: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder() // simulate an http response

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

func TestRegisterUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if it doesnt register the user", func(t *testing.T) {
		payload := models.Payload{
			Name:     "testing",
			LastName: "testingtesting",
			Email:    "valid@mail.com",
			Password: "contrasenia",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("failed to marshal payload: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Errorf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code: %d, got: %d", http.StatusCreated, rr.Code)
		}
	})
}
