package customerimporter

import (
	"fmt"
	"io"
	"net/mail"
	"sort"
	"strings"
)

// CountCustomerByDomain gets list of customer objects, counts for each domain how many
// customers it has and sorts it alphabetically
func CountCustomerByDomain(customers []Customer) (EmailDomainCustomerCountList, error) {
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

func CountCustomerByDomainFromCSV(r io.Reader) ([]EmailDomainCustomerCount, error) {
	customers, err := getCustomers(r)
	if err != nil {
		return nil, fmt.Errorf("error when processing CSV: %w", err)
	}
	return CountCustomerByDomain(customers)

}

// getDomainFromEmail retrieves domain for given email
func getDomainFromEmail(email string) (string, error) {
	address, err := mail.ParseAddress(email)

	if err != nil {
		return "", fmt.Errorf("invalid email: %q, %w", email, err)
	}

	emailParts := strings.Split(address.Address, "@")

	// parsing email through mail.ParseAddress should not let this happen, bu for sanity I'll check it
	if len(emailParts) != 2 {
		return "", fmt.Errorf("could not get domain from email: %q", email)
	}
	domain := emailParts[1]

	return domain, nil
}
