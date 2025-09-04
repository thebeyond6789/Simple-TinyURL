package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	
	"backend/handlers"
)

func main() {
	router := mux.NewRouter()
	
	// Test route
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test route works!")
	}).Methods("GET")
	
	// API routes
	router.HandleFunc("/api/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{slug}", handlers.RedirectURL).Methods("GET")
	
	// Serve static files (HTML, CSS, JS) - đặt sau các route API
	// Sử dụng đường dẫn tuyệt đối
	frontendPath, _ := filepath.Abs("../frontend")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(frontendPath)))
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}