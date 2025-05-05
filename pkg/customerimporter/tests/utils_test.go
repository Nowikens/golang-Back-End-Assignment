package customerimporter_test

import (
	"bytes"
	"encoding/csv"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

var silentLogger = slog.New(
	slog.NewTextHandler(io.Discard, nil),
)

// getTooFewColumnsCSVData returns io.reader with too few columns
func getTooFewColumnsCSVData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		{
			"first_name", "last_name", "email",
		},
		{
			"some_name", "some_last_name", "example@exmaple.com",
		},
	}

	return getCSVData(t, records)
}

// getTooManyColumnsCSVData returns io.reader with too many columns
func getTooManyColumnsCSVData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		{
			"first_name", "last_name", "email", "gender", "ip_address", "something", "else",
		},
	}

	return getCSVData(t, records)
}

// getCSVData generates inmemory csv from provided records to pass to functions
func getCSVData(t *testing.T, records [][]string) io.Reader {
	t.Helper()

	var buf bytes.Buffer
	csvWriter := csv.NewWriter(&buf)
	err := csvWriter.WriteAll(records)
	require.NoError(t, err, "Error when writing records")
	csvWriter.Flush()
	require.NoError(t, csvWriter.Error(), "Error when flushing records")
	return &buf
}

func getProperColumns() []string {
	return []string{
		"first_name", "last_name", "email", "gender", "ip_address",
	}
}

// getOneBadEmailRow returns io.reader with bad email in one row
func getOneBadEmailRow(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		getProperColumns(),
		{
			"first_name", "last_name", "email", "gender", "ip_address",
		},
		{
			"Mildred", "Hernandez", "bad_email", "Female", "38.194.51.128",
		},
		{
			"Mildred", "Hernandez", "example@example.com", "Female", "38.194.51.138",
		},
	}

	return getCSVData(t, records)
}