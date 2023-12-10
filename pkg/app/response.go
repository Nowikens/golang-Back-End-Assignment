package app

import (
	"github.com/google/uuid"
	"github.com/nowikens/customer_importer/pkg/customer"
)

type ProcessCSVResponse struct {
	ID         uuid.UUID `json:"id"`
	ResultLink string    `json:"result_link"`
}

type CSVResultResponse struct {
	Data customer.EmailDomainCustomerCountList `json:"data"`
}
