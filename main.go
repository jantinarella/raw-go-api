package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jantinarella/raw-go-api/router"
)

func main() {
	fmt.Println("Hello, World!")
	mux := router.NewRouter()
	port := ":8080"

	router := router.LoggingMiddleware(mux)

	log.Println("Server is starting on port", port, "...")
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("serverl faild:", err)
	}
}
