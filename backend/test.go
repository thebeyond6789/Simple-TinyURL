package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	// Test route
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test route works!")
	}).Methods("GET")
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}