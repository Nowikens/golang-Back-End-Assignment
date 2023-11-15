package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nowikens/customer_importer/pkg/api/response"
)

// HandleEmails accepts csv file, starts processing it and returns process ID
func HandleEmails(w http.ResponseWriter, r *http.Request) {

	response_data := response.FileResponse{
		ID: 1,
	}

	jsonResponse, err := json.Marshal(response_data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
