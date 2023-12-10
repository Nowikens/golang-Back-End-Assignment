package customer

import (
	"log/slog"
)

type CustomerService interface{}

type customerService struct {
	logger *slog.Logger
}

func NewCustomerService(logger *slog.Logger) *customerService {
	return &customerService{
		logger: logger,
	}
}
