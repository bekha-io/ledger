package services

import "github.com/bekha-io/ledger/domain/entities"

type JournalServiceI interface {
	// CreateJournal validates the entries and creates a journal with all the given entries
	CreateJournal(entries []*entities.Transaction) (*entities.Entry, error)
}
