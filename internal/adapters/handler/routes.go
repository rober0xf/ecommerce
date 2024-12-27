package handler

import (
	"ecommerce/internal/core/models"
	"ecommerce/internal/core/services"
	"ecommerce/pkg"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

var ValidateStruct = validator.New()

type Handler struct {
	Store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		pkg.WriteError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var payload models.PayloadLogin
	if err := pkg.ParseJSON(&payload, r); err != nil {
		pkg.WriteError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// validate json
	if err := ValidateStruct.Struct(payload); err != nil {
		fmt.Printf("validation error: %v\n", err)
		error_msg := err.(validator.ValidationErrors)
		pkg.WriteError(w, fmt.Sprintf("invalid input: %v", error_msg), http.StatusBadRequest)
		return
	}

	// check if the user exists
	user, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		pkg.WriteError(w, "user not found", http.StatusBadRequest)
		return
	}

	if !services.CheckPassword(user.Password, payload.Password) {
		pkg.WriteError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	pkg.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		pkg.WriteError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var payload models.PayloadRegister
	if err := pkg.ParseJSON(&payload, r); err != nil {
		pkg.WriteError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// validate json
	if err := ValidateStruct.Struct(payload); err != nil {
		fmt.Printf("validation error: %v\n", err)
		error_msg := err.(validator.ValidationErrors)
		pkg.WriteError(w, fmt.Sprintf("invalid input: %v", error_msg), http.StatusBadRequest)
		return
	}

	// check if the user exists
	_, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		pkg.WriteError(w, fmt.Sprintf("user %s already exists", payload.Email), http.StatusBadRequest)
		return
	}

	// hash password
	hashed_password, err := services.HashPassword(payload.Password)
	if err != nil {
		pkg.WriteError(w, "password error", http.StatusInternalServerError)
		return
	}

	// if it doesnt exists
	err = h.Store.CreateUser(&models.User{
		Name:     payload.Name,
		LastName: payload.LastName,
		Email:    payload.Email,
		Password: hashed_password,
	})

	if err != nil {
		pkg.WriteError(w, "error creating user", http.StatusInternalServerError)
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, `user created`)
}
