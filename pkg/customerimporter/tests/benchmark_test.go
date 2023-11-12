package customerimporter_test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/stretchr/testify/require"
)

func BenchmarkCountCustomerByDomainFromCSV(b *testing.B) {
	sizes := []int{100, 1000, 10_000, 100_000, 1000_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size: %d", size), func(b *testing.B) {
			reader := generateCSVReader(b, size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				customerimporter.CountCustomerByDomainFromCSV(reader)
			}
		})
	}
}
func BenchmarkCountCustomerByDomainFromCSVExampleData(b *testing.B) {

	b.Run("Example data", func(b *testing.B) {
		reader, err := os.Open("data/customers.csv")
		require.NoError(b, err, "Error during opening example data")

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			customerimporter.CountCustomerByDomainFromCSV(reader)
		}
	})
}

func generateCSVReader(b *testing.B, size int) io.Reader {
	b.Helper()
	records := [][]string{}
	// Generate CSV data with 5 columns
	for i := 0; i < size; i++ {
		records = append(
			records,
			[]string{"first_name"},
			[]string{fmt.Sprintf("last_name%d", i)},
			[]string{fmt.Sprintf("email.x%d.com", i)},
			[]string{"gender"},
			[]string{"127.0.0.1"},
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
