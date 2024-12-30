package handler

import (
	"ecommerce/internal/adapters/repository"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool // connection with the database
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	userStore := repository.NewUserStore(s.db)
	userHandler := NewUserHandler(userStore)
	userHandler.RegisterUserRoutes(subrouter)

	productStore := repository.NewProductStore(s.db)
	productHandler := NewProductHandler(productStore)
	productHandler.RegisterProductRoutes(subrouter)

	fmt.Print("running on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
