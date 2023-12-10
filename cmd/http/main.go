package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nowikens/customer_importer/pkg/app"
	"github.com/nowikens/customer_importer/pkg/customer"
)

func init() {
	p, _ := os.Getwd()
	fmt.Println(p)
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
func main() {
	port, portString := getPortInfo()

	logger := slog.Default()
	cs := customer.NewCustomerService(logger)
	a, err := app.New(
		cs,
	)
	if err != nil {
		panic(err)
	}
	logger.Info("Starting server", slog.Int("port", port))
	http.ListenAndServe(portString, a.GetRouter())
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
