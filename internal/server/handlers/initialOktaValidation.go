package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func VerificationHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET

	// Get the verification challenge from the header
	challenge := r.Header.Get("x-okta-verification-challenge")
	if challenge == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error("Verification: Missing x-okta-verification-challenge header")
		return
	}

	// Create the response payload
	response := map[string]string{"verification": challenge}

	// Marshal the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Verification: Failed to marshal JSON response", "error", err)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		slog.Error("Verification: Failed to write response", "error", err)
		return
	}

	slog.Info("Verification: Successfully responded to verification challenge", "challenge", challenge)
}
