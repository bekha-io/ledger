package memory

import (
	"context"
	"sync"

	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/errs"
	"github.com/bekha-io/ledger/domain/repository"
)

var _ repository.EntryRepositoryI = (*MemoryEntryRepository)(nil)

type MemoryEntryRepository struct {
	counter int
	entries sync.Map
}

// GetEntryByID implements repository.EntryRepositoryI.
func (m *MemoryEntryRepository) GetEntryByID(ctx context.Context, id uint) (*entities.Entry, error) {
	e, ok := m.entries.Load(id)
	if !ok {
		return nil, errs.ErrRecordNotFound
	}

	entry, ok := e.(*entities.Entry)
	if !ok {
		return nil, errs.ErrRecordNotFound
	}

	return entry, nil
}

// SaveEntry implements repository.EntryRepositoryI.
func (m *MemoryEntryRepository) SaveEntry(ctx context.Context, entry *entities.Entry) error {
	m.counter += 1
	entry.ID = uint(m.counter)
	m.entries.Store(m.counter, entry)
	return nil
}
