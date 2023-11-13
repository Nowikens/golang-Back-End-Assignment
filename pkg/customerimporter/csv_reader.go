package customerimporter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"strings"

	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

const (
	ProperColumnsString = "first_name,last_name,email,gender,ip_address"
)

var (
	ErrBadRow     = errors.New("bad row")
	ErrBadEmail   = errors.New("bad email")

	ProperColumns = strings.Split(ProperColumnsString, ",")

	ExpectedColumns = fmt.Sprintf("expected columns: %q", ProperColumnsString)
)

// getCustomers processes csv file, validates columns rows and emails, and creates list of Customer objects
func getCustomers(a *app.App, r io.Reader) ([]Customer, error) {
	customers := []Customer{}

	scanner := bufio.NewScanner(r)

	// processing customers data
	for scanner.Scan() {
		row := scanner.Text()

		rowData, err := getRowData(row)
		if err != nil {
			a.Logger.Warn("", err)
			// when structure is wrong there is not much we can do
			return nil, err
		}

		email, err := getEmailFromRow(rowData)
		if err != nil {
			a.Logger.Warn("", err)
			// when couldn't process email, log error and proceed to next row
			continue
		}

		customers = append(customers, Customer{email})
	}
	return customers, nil
}

// getRowData validates row, and returns splited row data
func getRowData(rowString string) ([]string, error) {
	rowData := strings.Split(rowString, ",")

	if len(rowData) != len(ProperColumns) {
		return nil, fmt.Errorf("%w: %s. Got row: %q", ErrBadRow, ExpectedColumns, rowString)
	}

	return rowData, nil
}

// getEmailFromRow takes splitted row data, searches for email value, validates it and returns it
func getEmailFromRow(rowData []string) (string, error) {
	email := rowData[2]

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", fmt.Errorf("%w: %q %w", ErrBadEmail, email, err)
	}

	return email, nil
}
