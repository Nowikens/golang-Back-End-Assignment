package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nowikens/customer_importer/pkg/api/server"
	"github.com/nowikens/customer_importer/pkg/app"
)

func init() {
	p, _ := os.Getwd()
	fmt.Println(p)
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
func main() {
	a := app.App{
		Logger: slog.Default(),
	}
	s := server.New()
	s.MountHandlers()

	port, portString := getPortInfo()
	a.Logger.Info("Starting server", slog.Int("port", port))
	http.ListenAndServe(portString, s.Router)
}
func getPortInfo() (port int, portString string) {
	portEnv, ok := os.LookupEnv("APP_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("APP_PORT not specified in environment variables")
	}
	portString = fmt.Sprintf(":%d", port)
	return
}