package customerimporter_test

import (
	"io"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
	"github.com/stretchr/testify/assert"
)

func TestCountCustomerByDomain(t *testing.T) {
	a := app.App{
		Logger: silentLogger,
	}
	testCases := []struct {
		desc   string
		input  []customerimporter.Customer
		output customerimporter.EmailDomainCustomerCountList
		err    error
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
		{
			desc: "bad email",
			input: []customerimporter.Customer{
				{
					Email: "abcds",
				},
			},
			err: customerimporter.ErrInvalidEmail,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := customerimporter.CountCustomerByDomain(&a, tC.input)
			if tC.err != nil {
				assert.ErrorIs(t, err, tC.err)
			}
			if tC.output != nil {
				assert.Equal(t, tC.output, data, "Data should be sorted and properly counted")
			}
			if tC.err == nil {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCountCustomerByDomainFromCSV(t *testing.T) {
	a := app.App{
		Logger: silentLogger,
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
			desc: "column order doesn't matter",
			csvReader: getCSVData(
				t,
				[][]string{
					{
						"first_name", "last_name", "gender", "ip_address", "email",
					},
					{
						"a", "b", "Female", "127.0.0.1", "a@x.com",
					},
					{
						"c", "d", "Female", "127.0.0.1", "a@y.com",
					},
					{
						"e", "f", "Female", "127.0.0.1", "a@x.com",
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
			desc: "'email' header required - headers row exist",
			csvReader: getCSVData(
				t,
				[][]string{
					{
						"first_name", "last_name", "gender", "ip_address", "not_an_email",
					},
					{
						"a", "b", "Female", "127.0.0.1", "a@x.com",
					},
				},
			),
			err: customerimporter.ErrEmailColumnNotFound,
		},
		{
			desc: "'email' header required - headers row not exist",
			csvReader: getCSVData(
				t,
				[][]string{
					{
						"a", "b", "Female", "127.0.0.1", "a@x.com",
					},
				},
			),
			err: customerimporter.ErrEmailColumnNotFound,
		},
		{
			desc: "empty file returns error - no headers",
			csvReader: getCSVData(
				t,
				[][]string{
					{},
				},
			),
			err: customerimporter.ErrBadHeaders,
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
		{
			desc:      "one bad email doesn't prevent finishing rest",
			csvReader: getOneBadEmailRow(t),
			err:       nil,
			output: customerimporter.EmailDomainCustomerCountList{
				{
					EmailDomain:   "example.com",
					CustomerCount: 1,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := customerimporter.CountCustomerByDomainFromCSV(&a, tC.csvReader)
			if tC.err != nil {
				assert.ErrorIs(t, err, tC.err)
			}
			if tC.err == nil {
				assert.NoError(t, err)
			}
			if tC.output != nil {
				assert.Equal(t, tC.output, result)
			}
		})
	}
}
