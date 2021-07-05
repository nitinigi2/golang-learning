package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nitinigi2/go-learning/testing/service"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := service.GetBooks()
	json.NewEncoder(w).Encode(books)
}
