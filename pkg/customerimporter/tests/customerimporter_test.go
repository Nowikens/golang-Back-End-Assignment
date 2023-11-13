package customerimporter_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
	"github.com/stretchr/testify/assert"
)

func TestCountCustomerByDomain(t *testing.T) {
	a := app.App{
		Logger: slog.Default(),
	}
	testCases := []struct {
		desc   string
		input  []customerimporter.Customer
		output customerimporter.EmailDomainCustomerCountList
	}{
		{
			desc: "happy case #1",
			input: []customerimporter.Customer{
				{
					Email: "c@z.com",
				},
				{
					Email: "a@x.com",
				},
				{
					Email: "b@y.com",
				},
				{
					Email: "d@x.com",
				},
			},
			output: customerimporter.EmailDomainCustomerCountList{
				{
					EmailDomain:   "x.com",
					CustomerCount: 2,
				},
				{
					EmailDomain:   "y.com",
					CustomerCount: 1,
				},
				{
					EmailDomain:   "z.com",
					CustomerCount: 1,
				},
			},
		},
		{
			desc: "happy case #2",
			input: []customerimporter.Customer{
				{
					Email: "b@y.com",
				},
				{
					Email: "a@x.com",
				},
				{
					Email: "d@x.com",
				},
			},
			output: customerimporter.EmailDomainCustomerCountList{
				{
					EmailDomain:   "x.com",
					CustomerCount: 2,
				},
				{
					EmailDomain:   "y.com",
					CustomerCount: 1,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := customerimporter.CountCustomerByDomain(&a, tC.input)
			assert.NoError(t, err, "CountCustomerByDomain should return no errors")

			assert.Equal(t, tC.output, data, "Data should be sorted and properly counted")
		})
	}

}

func TestCountCustomerByDomainFromCSV(t *testing.T) {
	a := app.App{
		Logger: slog.Default(),
	}
	testCases := []struct {
		desc      string
		csvReader io.Reader
		output    customerimporter.EmailDomainCustomerCountList
		err       error
	}{
		{
			desc: "happy case",
			csvReader: getCSVData(
				t,
				[][]string{
					getProperColumns(),
					{
						"a", "b", "a@x.com", "Female", "127.0.0.1",
					},
					{
						"c", "d", "a@y.com", "Female", "127.0.0.1",
					},
					{
						"e", "f", "a@x.com", "Female", "127.0.0.1",
					},
				},
			),
			output: customerimporter.EmailDomainCustomerCountList{
				{
					EmailDomain:   "x.com",
					CustomerCount: 2,
				},
				{
					EmailDomain:   "y.com",
					CustomerCount: 1,
				},
			},
		},
		{
			desc:      "too few columns",
			csvReader: getTooFewColumnsCSVData(t),
			err:       customerimporter.ErrBadRow,
		},
		{
			desc:      "too many columns",
			csvReader: getTooManyColumnsCSVData(t),
			err:       customerimporter.ErrBadRow,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := customerimporter.CountCustomerByDomainFromCSV(&a, tC.csvReader)
			if tC.err != nil {
				assert.ErrorIs(t, err, tC.err)
			}
			if tC.output != nil {
				assert.Equal(t, tC.output, result)
			}
		})
	}
}
