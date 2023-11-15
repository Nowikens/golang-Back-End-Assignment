package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/nowikens/customer_importer/pkg/api/server"
	"github.com/nowikens/customer_importer/pkg/app"
)

func main() {

	a := app.App{
		Logger: slog.Default(),
	}
	s := server.New()
	s.MountHandlers()

	port := 8080
	portString := fmt.Sprintf(":%d", port)
	a.Logger.Info("Starting server", slog.Int("port", port))
	http.ListenAndServe(portString, s.Router)
}
