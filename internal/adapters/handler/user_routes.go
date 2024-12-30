package handler

import (
	"ecommerce/config"
	"ecommerce/internal/core/models"
	"ecommerce/internal/core/services"
	"ecommerce/pkg"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var ValidateStruct = validator.New()

type UserHandler struct {
	Store models.UserStore
}

func NewUserHandler(store models.UserStore) *UserHandler {
	return &UserHandler{Store: store}
}

func (h *UserHandler) RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
}

func (h *UserHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		log.Printf("error: %v", err)
		pkg.WriteError(w, "user not found", http.StatusBadRequest)
		return
	}

	if err := services.CheckPassword(user.Password, payload.Password); err != nil {
		pkg.WriteError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	jwtConfig := config.LoadJWTConfig() // load jwt config
	if jwtConfig == nil {
		pkg.WriteError(w, "jwt not found", http.StatusInternalServerError)
		return
	}

	token, err := services.CreateJWT(*jwtConfig, user.ID) // returns the token as a string
	if err != nil {
		pkg.WriteError(w, "unable to create token", http.StatusInternalServerError)
	}

	pkg.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("error: %v", err)
		pkg.WriteError(w, "error creating user", http.StatusInternalServerError)
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, `user created`)
}
