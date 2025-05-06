package customerimporter_test

import (
	"encoding/json"
	"testing"

	"github.com/nowikens/customer_importer/pkg/customerimporter"
	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
	"github.com/stretchr/testify/assert"
)

func TestExportToJson(t *testing.T) {
	a := app.App{
		Logger: silentLogger,
	}
	testCases := []struct {
		desc         string
		input        customerimporter.EmailDomainCustomerCountList
		expectedJson string
		err          error
	}{
		{
			desc: "happy case",
			input: customerimporter.EmailDomainCustomerCountList{
				{
					EmailDomain:   "x.com",
					CustomerCount: 2,
				},
				{
					EmailDomain:   "y.com",
					CustomerCount: 1,
				},
			},
			expectedJson: `[{"email_domain":"x.com","customer_count":2},{"email_domain":"y.com","customer_count":1}]`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// run function
			data, err := customerimporter.ExportToJson(&a, tC.input)
			assert.NoError(t, err)

			// unmarshall expected and acual results
			// expected
			var ExpectedEmailDomainCustomerCountList customerimporter.EmailDomainCustomerCountList
			err = json.Unmarshal([]byte(tC.expectedJson), &ExpectedEmailDomainCustomerCountList)
			assert.NoError(t, err, "Error during unmarshaling expected JSON")
			// actual
			var resultEmailDomainCustomerCountList customerimporter.EmailDomainCustomerCountList
			err = json.Unmarshal(data, &resultEmailDomainCustomerCountList)
			assert.NoError(t, err, "Error during unmarshaling actual JSON")

			// assertions
			if tC.err != nil {
				assert.ErrorIs(t, err, tC.err)
			}
			if tC.expectedJson != "" {
				assert.Equal(t, ExpectedEmailDomainCustomerCountList, resultEmailDomainCustomerCountList)
			}
			if tC.err == nil {
				assert.NoError(t, err)
			}
		})
	}
}
