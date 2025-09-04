package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Slug string `json:"slug"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate URL
	if req.URL == "" {
		sendErrorResponse(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Generate unique slug
	slug := utils.GenerateSlug()
	// In a real application, you would check if the slug already exists and regenerate if needed
	// For this simple example, we'll assume it's unique

	// Save mapping
	models.SetURL(slug, req.URL)

	// Return response
	response := ShortenResponse{Slug: slug}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:] // Remove leading slash

	// Get original URL
	originalURL, exists := models.GetURL(slug)
	if !exists {
		http.NotFound(w, r)
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := ErrorResponse{Error: message}
	json.NewEncoder(w).Encode(response)
}