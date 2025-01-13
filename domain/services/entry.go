package services

import (
	"context"
	"fmt"

	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/repository"
)

type EntryServiceI interface {
	// CreateEntry validates the entries and creates a journal with all the given entries
	CreateEntry(ctx context.Context, transactions []*entities.Transaction) (*entities.Entry, error)
}

type EntryService struct {
	Repo repository.RepoFacade
}

func (s *EntryService) CreateEntry(ctx context.Context, transactions []*entities.Transaction) (*entities.Entry, error) {
	// Collecting IDs in order to check if they exist
	for _, tr := range transactions {
		acc, err := s.Repo.Accounts.GetAccountByID(ctx, tr.AccountID)
		if err != nil || acc == nil {
			return nil, fmt.Errorf("invalid account id: %v account does not exist", tr.AccountID)
		}

		if err := acc.Process(tr); err != nil {
			return nil, fmt.Errorf("cannot settle transaction (%v) to account %v (%v)", tr.Description, acc.ID, acc.Name)
		}
	}

	entry := entities.NewEntry(transactions...)
	if err := entry.Validate(); err != nil {
		return nil, err
	}

	if err := s.Repo.Entries.SaveEntry(ctx, entry); err != nil {
		return nil, err
	}

	return entry, nil
}
