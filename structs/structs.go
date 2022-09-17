package structs

import (
	constants "github.com/osoriojuan/go-integration-grcp/constants"
)

type DefaultRequest struct {
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
}

type DefaultResponse struct {
	Result  string `json:"operation"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ExchangeRateRequest struct {
	ID                   string `json:"id"`
	CurrenciesComparison constants.CurrencyType
}
