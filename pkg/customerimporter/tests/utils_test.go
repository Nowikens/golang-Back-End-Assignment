package customerimporter_test

import (
	"bytes"
	"encoding/csv"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// getTooFewColumnsCSVData returns io.reader with too few columns
func getTooFewColumnsCSVData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		{
			"first_name", "last_name", "ip_address",
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

// getBadColumnsCSVData returns io.reader with bad columns
func getBadColumnsCSVData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		{
			"first_name", "badColumn", "email", "gender", "ip_address",
		},
	}

	return getCSVData(t, records)
}

// getBadRowData returns io.reader with one corrupted row
func getBadRowData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		getProperColumns(),
		{
			"Mildred", "Hernandez", "Female", "38.194.51.128",
		},
	}

	return getCSVData(t, records)
}

// getBadEmailCSVData returns properly structured data but with wrong email
func getBadEmailCSVData(t *testing.T) io.Reader {
	t.Helper()
	records := [][]string{
		getProperColumns(),
		{
			"Mildred", "Hernandez", "badEmail", "Female", "38.194.51.128",
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
