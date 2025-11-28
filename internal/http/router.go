package http 

import (
	"net/http"

	"template/internal/http/handlers"
)


func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	// mux.HandleFunc("/users", handlers.UserHandler)
	// mux.HandleFunc("/products", handlers.ProductHandler)

	return mux
}