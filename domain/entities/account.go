package entities

import (
	"fmt"
	"slices"

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

func (a *Account) Validate() error {
	if !slices.Contains(types.AccountTypes, a.Type) {
		return fmt.Errorf("invalid account type: %v", a.Type)
	}
	return nil
}

func (a *Account) Process(tr *Transaction) error {
	if tr.Type == types.TransactionTypeDebit {
		a.Balance = a.Balance.Add(tr.Amount)
	} else if tr.Type == types.TransactionTypeCredit {
		a.Balance = a.Balance.Sub(tr.Amount)
	} else {
		return fmt.Errorf("invalid transaction type: %v", tr.Type)
	}
	return nil
}
