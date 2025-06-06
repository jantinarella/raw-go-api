package router

import (
	"net/http"

	"github.com/jantinarella/raw-go-api/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/users", handler.UserHandler)
	return mux
}
