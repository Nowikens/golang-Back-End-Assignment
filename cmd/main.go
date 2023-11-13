package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

func main() {
	filePath := flag.String("file", "customers.csv", "File with customers data to process")
	flag.Parse()
	a := app.App{
		Logger: slog.Default(),
	}

	file, err := os.Open(*filePath)
	if err != nil {
		a.Logger.Error(
			"error during opening the file",
			slog.String("path", *filePath),
		)
		return
	}
	defer file.Close()

	result, err := customerimporter.CountCustomerByDomainFromCSV(&a, file)
	if err != nil {
		a.Logger.Error("err during reading csv")
		return
	}
	fmt.Printf("%+v", result)

}
