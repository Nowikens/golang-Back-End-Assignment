package customerimporter

import (
	"errors"
	"fmt"
	"io"
	"net/mail"
	"sort"
	"strings"

	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

var (
	ErrInvalidEmail        = errors.New("invalid email")
	ErrNotEnoughEmailParts = errors.New("could not get domain from email")
)

// CountCustomerByDomain gets list of customer objects, counts for each domain how many
// customers it has and sorts it alphabetically
func CountCustomerByDomain(app *app.App, customers []Customer) (EmailDomainCustomerCountList, error) {
	result := make(EmailDomainCustomerCountList, 0)

	m := make(map[string]int, 0)

	// count
	for _, customer := range customers {
		emailDomain, err := getDomainFromEmail(customer.Email)
		if err != nil {
			return nil, err
		}
		m[emailDomain]++
	}

	for k := range m {
		result = append(
			result,
			EmailDomainCustomerCount{
				EmailDomain:   k,
				CustomerCount: m[k],
			},
		)
	}

	sort.Sort(result)
	return result, nil
}

// CountCustomerByDomainFromCSV is a wrapper aroung CountCustomerByDomain designed specifically for CSV format.
func CountCustomerByDomainFromCSV(a *app.App, r io.Reader) (EmailDomainCustomerCountList, error) {
	customers, err := getCustomers(a, r)
	if err != nil {
		return nil, fmt.Errorf("error when processing CSV: %w", err)
	}
	return CountCustomerByDomain(a, customers)

}

// getDomainFromEmail retrieves domain for given email
func getDomainFromEmail(email string) (string, error) {
	address, err := mail.ParseAddress(email)

	if err != nil {
		return "", fmt.Errorf("%w: %q, %w", ErrInvalidEmail, email, err)
	}

	emailParts := strings.Split(address.Address, "@")

	// parsing email through mail.ParseAddress should not let this happen, but for sanity I'll check it
	if len(emailParts) != 2 {
		return "", fmt.Errorf("%w: %q, %w", ErrNotEnoughEmailParts, email, err)
	}
	domain := emailParts[1]

	return domain, nil
}
