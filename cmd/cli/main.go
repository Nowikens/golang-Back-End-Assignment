package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nowikens/customer_importer/pkg/customer"
)

func main() {
	filePath := flag.String("file", "customers.csv", "File with customers data to process")
	flag.Parse()
	logger := slog.Default()
	s := customer.NewCustomerService(logger)

	file, err := os.Open(*filePath)
	if err != nil {
		logger.Error(
			"error during opening the file",
			slog.String("path", *filePath),
		)
		return
	}
	defer file.Close()

	result, err := s.CountCustomerByDomainFromCSV(file)
	if err != nil {
		logger.Error("err during reading csv")
		return
	}
	fmt.Printf("%+v", result)

}
