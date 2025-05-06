package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

const (
	unset = ""
	json  = "json"
)

func main() {
	// wrap in run to avoid handling mixin os.Exit and defers
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run() error {
	// handle flags
	filePath := flag.String("file", unset, "File with customers data to process.")
	outputFormat := flag.String("format", json, "Output format. Valid options: json.")
	outputPath := flag.String("output", unset, "Output path. If not provided - output will be printed to the console.")
	flag.Parse()
	// validate flag options
	if *outputFormat != json {
		return fmt.Errorf("unsupported format: %s", *outputFormat)
	}
	if *filePath == unset {
		return errors.New("error: --output flag is required")
	}

	// handle cusomerimporter app execution
	a := app.App{
		Logger: slog.Default(),
	}
	file, err := os.Open(*filePath)
	if err != nil {
		return fmt.Errorf("open file to read: %w", err)
	}
	defer file.Close()
	result, err := customerimporter.CountCustomerByDomainFromCSV(&a, file)
	if err != nil {
		return fmt.Errorf("read csv: %w", err)
	}

	// handle output
	var fileData []byte

	// handle output format
	switch *outputFormat {
	case json:
		jsonData, err := customerimporter.ExportToJson(&a, result)
		if err != nil {
			return fmt.Errorf("export to json: %w", err)
		}
		fileData = jsonData
	default:
		return fmt.Errorf("unsupported format: %s", *outputFormat)
	}

	// handle output target
	if *outputPath == unset {
		// print to console
		fmt.Println(string(fileData))
	} else {
		// write to file
		file, err := os.OpenFile(*outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("open file to write: %w", err)
		}
		_, err = file.Write(fileData)
		if err != nil {
			return fmt.Errorf("write to file: %w", err)
		}
		defer file.Close()
	}
	return nil
}
