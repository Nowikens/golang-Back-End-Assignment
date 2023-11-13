package customerimporter_test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
	"github.com/stretchr/testify/require"
)

func BenchmarkCountCustomerByDomainFromCSV(b *testing.B) {
	sizes := []int{100, 1000, 10_000, 10_000, 1000_000}
	a := app.App{
		// turn off logging so it doesn't obscure benchmark's output
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size: %d", size), func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				reader := generateCSVReader(b, size)
				b.StartTimer()
				_, err := customerimporter.CountCustomerByDomainFromCSV(&a, reader)
				b.StopTimer()
				require.NoError(b, err)
			}
		})
	}
}
func BenchmarkCountCustomerByDomainFromCSVExampleData(b *testing.B) {
	a := app.App{
		// turn off logging so it doesn't obscure benchmark's output
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	b.Run("Example data", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reader, err := os.Open("data/customers.csv")
			require.NoError(b, err, "Error during opening example data")
			b.StartTimer()
			_, err = customerimporter.CountCustomerByDomainFromCSV(&a, reader)
			b.StopTimer()
			require.NoError(b, err)
		}
	})
}

// generateCSVReader generates io.reader with dummy data with specified length
func generateCSVReader(b *testing.B, size int) io.Reader {
	b.Helper()
	records := [][]string{}
	records = append(records, getProperColumns())
	for i := 0; i < size; i++ {
		records = append(
			records,
			[]string{"x", fmt.Sprintf("y%d", i), fmt.Sprintf("x@example%d.com", i), "gender", "127.0.0.1"},
		)
	}

	var buf bytes.Buffer
	csvWriter := csv.NewWriter(&buf)
	err := csvWriter.WriteAll(records)
	require.NoError(b, err, "Error when writing records")
	csvWriter.Flush()
	require.NoError(b, csvWriter.Error(), "Error when flushing records")
	return &buf
}
