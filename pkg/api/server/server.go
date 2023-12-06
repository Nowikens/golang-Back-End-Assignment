package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nowikens/customer_importer/pkg/customer/handlers"
)

type server struct {
	Router *chi.Mux
}

func New() *server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return &server{
		Router: r,
	}
}

// MountHandlers mounts all handlers
func (s *server) MountHandlers() {
	s.Router.Post(Emails, handlers.HandleEmails)
}
