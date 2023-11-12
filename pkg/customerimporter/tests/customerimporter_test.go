package customerimporter_test

import (
	"io"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/stretchr/testify/require"
)

func TestCountCustomerByDomain(t *testing.T) {
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
			data, err := customerimporter.CountCustomerByDomain(tC.input)
			require.NoError(t, err, "CountCustomerByDomain should return no errors")

			require.Equal(t, tC.output, data, "Data should be sorted and properly counted")

		})
	}

}

func TestCountCustomerByDomainFromCSV(t *testing.T) {
	testCases := []struct {
		desc      string
		csvReader io.Reader
		err       error
	}{
		{
			desc:      "too few columns",
			csvReader: getTooFewColumnsCSVData(t),
			err:       customerimporter.ErrBadColumns,
		},
		{
			desc:      "too many columns",
			csvReader: getTooManyColumnsCSVData(t),
			err:       customerimporter.ErrBadColumns,
		},
		{
			desc:      "bad columns",
			csvReader: getBadColumnsCSVData(t),
			err:       customerimporter.ErrBadColumns,
		},
		{
			desc:      "bad row data",
			csvReader: getBadRowData(t),
			err:       customerimporter.ErrBadRow,
		},
		{
			desc:      "bad email",
			csvReader: getBadEmailCSVData(t),
			err:       customerimporter.ErrBadEmail,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := customerimporter.CountCustomerByDomainFromCSV(tC.csvReader)
			if tC.err != nil {
				require.ErrorIs(t, err, tC.err)
			}
		})
	}
}
