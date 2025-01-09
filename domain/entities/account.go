package entities

import (
	"github.com/bekha-io/ledger/domain/types"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        types.AccountType `json:"type"`
	Balance     decimal.Decimal   `json:"balance"`
}
