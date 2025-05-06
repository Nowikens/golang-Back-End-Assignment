package customerimporter

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/nowikens/customer_importer/pkg/customerimporter/app"
)

var (
	ErrMarshal = errors.New("marhsaling error")
)

// ExportToJson takes EmailDomainCustomerCountList data and converts it to JSON
func ExportToJson(a *app.App, customers EmailDomainCustomerCountList) ([]byte, error) {
	jsonData, err := json.Marshal(customers)
	if err != nil {
		a.Logger.Error(
			"Error while exporting to JSON",
			slog.Any("error", err),
		)
		return nil, err
	}
	return jsonData, nil
}
