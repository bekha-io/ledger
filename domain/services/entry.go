package services

import (
	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/repository"
)

type EntryServiceI interface {
	// CreateEntry validates the entries and creates a journal with all the given entries
	CreateEntry(transactions []*entities.Transaction) (*entities.Entry, error)
}

type EntryService struct {
	Repo repository.RepoFacade
}

func (s *EntryService) CreateEntry(transactions []*entities.Transaction) (*entities.Entry, error) {
	return nil, nil
}
