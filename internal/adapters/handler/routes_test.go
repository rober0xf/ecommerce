package handler

import (
	"bytes"
	"ecommerce/internal/core/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
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

	payload := models.Payload{
		Name:     "namee",
		LastName: "lastname",
		Email:    "valid@gmail.com",
		Password: "123434324",
	}

	marshalled, err := json.Marshal(payload) // convert the struct to a valid json
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	t.Run("should fail if the payload is invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder() // simulate an http response

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		// handleRegister returns an http bad request if is nil
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}
