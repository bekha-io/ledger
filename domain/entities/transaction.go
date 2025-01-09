package entities

import (
	"github.com/bekha-io/ledger/domain/types"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID          uint                  `json:"id"`
	AccountID   string                `json:"account_id"`
	Description string                `json:"description"`
	Type        types.TransactionType `json:"type"`
	Amount      decimal.Decimal       `json:"amount"`
}
