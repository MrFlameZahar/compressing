package ports

import (
	"bd/bd/internal/ports/http/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// Инициализация роутера

func NewRouter() http.Handler {

	authMux := mux.NewRouter()
	authMux.HandleFunc("/compress", handlers.Compress).Methods("POST")
	authMux.HandleFunc("/compress", handlers.GetImage).Methods("GET")

	return authMux
}
