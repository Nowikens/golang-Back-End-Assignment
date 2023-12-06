package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// HandleEmails accepts csv file, starts processing it and returns process ID
func HandleEmails(w http.ResponseWriter, r *http.Request) {

	response_data := FileResponse{
		ID: 1,
	}

	render.JSON(w, r, response_data)
	render.Status(r, http.StatusCreated)
}
