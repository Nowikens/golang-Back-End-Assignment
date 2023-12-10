package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nowikens/customer_importer/pkg/customer"
)

type App struct {
	router          *chi.Mux
	customerService customer.CustomerService
}

func New(cs customer.CustomerService) (*App, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return &App{
		customerService: cs,
		router:          r,
	}, nil

}

func (a *App) GetRouter() http.Handler {
	return a.router
}
