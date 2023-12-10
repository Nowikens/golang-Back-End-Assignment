package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nowikens/customer_importer/pkg/customer"
)

// MountHandlers mounts all handlers

// HandleEmails accepts csv file, starts processing it and returns process ID
func (a *App) HandleProcessCSV(w http.ResponseWriter, r *http.Request) {

	response_data := ProcessCSVResponse{
		ID: uuid.New(),
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response_data)
}

// HandleCSVResult returns processed csv data
func (a *App) HandleResult(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "id"))
	response_data := CSVResultResponse{
		Data: customer.EmailDomainCustomerCountList{},
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response_data)
}
