package handler

import (
	"ecommerce/internal/core/models"
	"ecommerce/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ProductHandler struct {
	Store models.ProductStore
}

func NewProductHandler(store models.ProductStore) *ProductHandler {
	return &ProductHandler{Store: store}
}

func (h *ProductHandler) RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *ProductHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Store.GetProducts()
	if err != nil {
		pkg.WriteError(w, "unable to fetch products", http.StatusInternalServerError)
		return
	}

	pkg.WriteJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	product, err := h.Store.CreateProduct()
	if err != nil {
		pkg.WriteError(w, "unable to create product", http.StatusInternalServerError)
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, product)
}
