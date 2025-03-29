package utility

// package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse is a structure to standardize error responses
type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"status"`
}

// Controller struct contains the response handling logic
type Controller struct {
	response *ResponseHandler
}

// ResponseHandler can contain common response handling methods
type ResponseHandler struct{}

// HTTPFail is a method that sends a structured error response
func (r *ResponseHandler) HTTPFail(w http.ResponseWriter, statusCode int, message string, details string) {
	// Set the appropriate HTTP status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Create the error response object
	errorResponse := ErrorResponse{
		Message: message,
		Details: details,
		Status:  statusCode,
	}

	// Write the JSON error response to the response writer
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		// If encoding the JSON fails, log the error and send a generic response
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
