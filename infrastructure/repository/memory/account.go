package memory

import (
	"context"
	"sync"

	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/errs"
	"github.com/bekha-io/ledger/domain/repository"
)

var _ repository.AccountRepositoryI = (*MemoryAccountRepository)(nil)

type MemoryAccountRepository struct {
	accounts sync.Map
}

// GetAccountByID implements repository.AccountRepositoryI.
func (m *MemoryAccountRepository) GetAccountByID(ctx context.Context, id string) (*entities.Account, error) {
	a, ok := m.accounts.Load(id)
	if !ok {
		return nil, errs.ErrRecordNotFound
	}

	account, ok := a.(*entities.Account)
	if !ok {
		return nil, errs.ErrRecordNotFound
	}

	return account, nil
}

// GetAllAccounts implements repository.AccountRepositoryI.
func (m *MemoryAccountRepository) GetAllAccounts(ctx context.Context) ([]*entities.Account, error) {
	var accounts []*entities.Account

	m.accounts.Range(func(k, v any) bool {
		account, ok := v.(*entities.Account)
		if ok {
			accounts = append(accounts, account)
		}
		return true
	})

	return accounts, nil
}

// SaveAccount implements repository.AccountRepositoryI.
func (m *MemoryAccountRepository) SaveAccount(ctx context.Context, account *entities.Account) error {
	m.accounts.Store(account.ID, account)
	return nil
}
