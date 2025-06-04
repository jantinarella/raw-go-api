package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("UserHandler"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "John Doe",
	}
	json.NewEncoder(w).Encode(user)
}
