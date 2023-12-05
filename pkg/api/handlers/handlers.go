package handlers

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/nowikens/customer_importer/pkg/api/response"
)

// HandleEmails accepts csv file, starts processing it and returns process ID
func HandleEmails(w http.ResponseWriter, r *http.Request) {

	response_data := response.FileResponse{
		ID: 1,
	}

	render.JSON(w, r, response_data)
	render.Status(r, http.StatusCreated)
}
