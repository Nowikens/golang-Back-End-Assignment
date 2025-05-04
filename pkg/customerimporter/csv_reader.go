package customerimporter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/mail"
	"strings"

	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

const (
	ProperColumnsString = "first_name,last_name,email,gender,ip_address"
)

var (
	ErrBadRow              = errors.New("bad row")
	ErrBadEmail            = errors.New("bad email")
	ErrBadHeaders          = errors.New("bad headers")
	ErrEmailColumnNotFound = errors.New("email column not found")

	ProperColumns = strings.Split(ProperColumnsString, ",")

	ExpectedColumns = fmt.Sprintf("expected columns: %q", ProperColumnsString)
)

// getCustomers processes csv file, validates columns rows and emails, and creates list of Customer objects
func getCustomers(a *app.App, r io.Reader) ([]Customer, error) {
	customers := []Customer{}

	reader := csv.NewReader(r)

	// reading first row which should be headers since we're not assuming email column position
	headers, err := reader.Read()
	if err != nil {
		a.Logger.Error(
			"Error during reading headers",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("%w: %w", ErrBadHeaders, err)
	}

	emailIndex, err := getEmailIndex(headers)
	if err != nil {
		a.Logger.Error(
			"Email column not found",
			slog.Any("error", err),
		)
		return nil, err
	}

	lineCounter := 0
	// processing customers data
	for {
		rowData, err := reader.Read()
		if err == io.EOF {
			break
		}
		lineCounter++

		err = validateRowData(rowData)
		if err != nil {
			a.Logger.Warn(
				"Error while reading row",
				slog.Int("line", lineCounter),
				slog.Any("error", err),
			)
			// when couldn't process row, log error and proceed to next row
			continue
		}

		email, err := getEmailFromRow(rowData, emailIndex)
		if err != nil {
			a.Logger.Warn(
				"Error while reading email",
				slog.Int("line", lineCounter),
				slog.Any("error ", err),
			)
			// when couldn't process email, log error and proceed to next row
			continue
		}

		customers = append(customers, Customer{email})
	}

	return customers, nil
}

// validateRowData validates row
func validateRowData(rowData []string) error {
	if len(rowData) != len(ProperColumns) {
		return fmt.Errorf("%w: %s. Got row: %q", ErrBadRow, ExpectedColumns, rowData)
	}

	return nil
}

// getEmailFromRow takes splitted row data,  searches for email value, validates it and returns it
func getEmailFromRow(rowData []string, emailPostition int) (string, error) {
	email := rowData[emailPostition]

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", fmt.Errorf("%w: %q %w", ErrBadEmail, email, err)
	}

	return email, nil
}

// getEmailIndex searches for email column index in headers
func getEmailIndex(headers []string) (int, error) {
	for i, h := range headers {
		if strings.EqualFold(h, "email") {
			return i, nil
		}
	}

	return -1, ErrEmailColumnNotFound
}
