package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jantinarella/raw-go-api/router"
)

func main() {
	mux := router.NewRouter()
	port := ":8080"

	router := router.LoggingMiddleware(mux)

	log.Println("Server is starting on port", port)
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  6 * time.Second,
	}
	log.Fatal(server.ListenAndServe())

}
