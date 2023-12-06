package customer_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/nowikens/customer_importer/pkg/app"
	"github.com/nowikens/customer_importer/pkg/customer"
	"github.com/stretchr/testify/assert"
)

func TestCountCustomerByDomain(t *testing.T) {
	a := app.App{
		Logger: slog.Default(),
	}
	testCases := []struct {
		desc   string
		input  []customer.Customer
		output customer.EmailDomainCustomerCountList
	}{
		{
			desc: "happy case #1",
			input: []customer.Customer{
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
			output: customer.EmailDomainCustomerCountList{
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
			input: []customer.Customer{
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
			output: customer.EmailDomainCustomerCountList{
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
			data, err := customer.CountCustomerByDomain(&a, tC.input)
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
		output    customer.EmailDomainCustomerCountList
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
			output: customer.EmailDomainCustomerCountList{
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
			desc:      "too few columns doesn't prevent finishing",
			csvReader: getTooFewColumnsCSVData(t),
			err:       nil,
		},
		{
			desc:      "too many columns doesn't prevent finishing",
			csvReader: getTooManyColumnsCSVData(t),
			err:       nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := customer.CountCustomerByDomainFromCSV(&a, tC.csvReader)
			if tC.err != nil {
				assert.ErrorIs(t, err, tC.err)
			}
			if tC.output != nil {
				assert.Equal(t, tC.output, result)
			}
		})
	}
}
