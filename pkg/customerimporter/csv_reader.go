package customerimporter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"strings"
)

const (
	ProperColumnsString = "first_name,last_name,email,gender,ip_address"
)

var (
	ErrBadColumns = errors.New("bad columns")
	ErrBadRow     = errors.New("bad row")
	ErrBadEmail   = errors.New("bad email")

	ProperColumns = strings.Split(ProperColumnsString, ",")

	ExpectedColumns = fmt.Sprintf("expected columns: %q", ProperColumnsString)
)

// getCustomers processes csv file, validates columns rows and emails, and creates list of Customer objects
func getCustomers(r io.Reader) ([]Customer, error) {
	customers := []Customer{}

	scanner := bufio.NewScanner(r)

	// ommit the first line with column names when processing
	scanner.Scan()
	columns := scanner.Text()
	if err := validateColumnsRow(columns); err != nil {
		return nil, err
	}

	// processing customers data
	for scanner.Scan() {
		row := scanner.Text()

		rowData, err := getRowData(row)
		if err != nil {
			return nil, err
		}

		email, err := getEmailFromRow(rowData)
		if err != nil {
			return nil, err
		}

		customers = append(customers, Customer{email})
	}
	return customers, nil
}

// validateColumnsRow validates first CSV's row with column names
func validateColumnsRow(columnsString string) error {
	if columnsString != ProperColumnsString {
		return fmt.Errorf("%w; %s, got %q", ErrBadColumns, ExpectedColumns, columnsString)
	}
	return nil
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
