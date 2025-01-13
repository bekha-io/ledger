package entities

import (
	"github.com/bekha-io/ledger/domain/errs"
	"github.com/bekha-io/ledger/domain/types"
	"github.com/shopspring/decimal"
)

// Entry is a collection of transactions that
type Entry struct {
	ID           uint               `json:"id"`
	Status       types.EntityStatus `json:"status"`
	Transactions []*Transaction     `json:"transactions"`
}

func NewEntry(trs ...*Transaction) *Entry {
	return &Entry{
		Status:       types.EntityStatusDraft,
		Transactions: trs,
	}
}

func (j *Entry) TotalDebit() decimal.Decimal {
	var total decimal.Decimal
	for _, tr := range j.Transactions {
		if tr.Type == types.TransactionTypeDebit {
			total = total.Add(tr.Amount)
		}
	}
	return total
}

func (j *Entry) TotalCredit() decimal.Decimal {
	var total decimal.Decimal
	for _, tr := range j.Transactions {
		if tr.Type == types.TransactionTypeCredit {
			total = total.Add(tr.Amount)
		}
	}
	return total
}

func (j *Entry) Validate() error {
	var totalBalance decimal.Decimal

	for _, tr := range j.Transactions {
		if tr.Type == types.TransactionTypeDebit {
			totalBalance = totalBalance.Add(tr.Amount)
			continue
		}
		totalBalance = totalBalance.Sub(tr.Amount)
	}

	if !totalBalance.IsZero() {
		return errs.ErrEntryBalanceMustBeZero
	}

	return nil
}
