package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user?name=Jane", nil)
	w := httptest.NewRecorder()
	getUser(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
