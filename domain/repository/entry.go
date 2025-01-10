package repository

import (
	"context"

	"github.com/bekha-io/ledger/domain/entities"
)

type EntryRepositoryI interface {
	GetEntryByID(ctx context.Context, id uint) (*entities.Entry, error)
	SaveEntry(ctx context.Context, entry *entities.Entry) error
}
