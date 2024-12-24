package handler

import (
	"ecommerce/internal/adapters/repository"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
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
    userStore := repository.NewStore(s.db)
    userHandler := NewHandler(userStore)
	userHandler.InitRoutes(subrouter)

	fmt.Print("running on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
