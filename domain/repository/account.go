package repository

import (
	"context"

	"github.com/bekha-io/ledger/domain/entities"
)

type AccountRepositoryI interface {
	GetAllAccounts(ctx context.Context) ([]*entities.Account, error)
	GetAccountByID(ctx context.Context, id string) (*entities.Account, error)
	SaveAccount(ctx context.Context, account *entities.Account) error
}
