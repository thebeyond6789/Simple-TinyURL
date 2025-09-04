package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test route works!")
	})
	
	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}