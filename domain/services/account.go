package services

import (
	"context"

	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/repository"
	"github.com/bekha-io/ledger/domain/types"
)

type CreateAccountIn struct {
	ID          string
	Name        string
	Description string
	Type        types.AccountType
}

type UpdateAccountIn struct {
	ID          string
	Name        string
	Description string
}

type AccountServiceI interface {
	// CreateAccount validates the account and creates a new account with the given details
	CreateAccount(ctx context.Context, in CreateAccountIn) (*entities.Account, error)
	UpdateAccount(ctx context.Context, in UpdateAccountIn) error
	GetAllAccounts(ctx context.Context) ([]*entities.Account, error)
	GetAccountByID(ctx context.Context, id string) (*entities.Account, error)
}

var _ AccountServiceI = (*AccountService)(nil)

type AccountService struct {
	Repo repository.RepoFacade
}

// CreateAccount implements AccountServiceI.
func (a *AccountService) CreateAccount(ctx context.Context, in CreateAccountIn) (*entities.Account, error) {
	panic("unimplemented")
}

// GetAccountByID implements AccountServiceI.
func (a *AccountService) GetAccountByID(ctx context.Context, id string) (*entities.Account, error) {
	panic("unimplemented")
}

// GetAllAccounts implements AccountServiceI.
func (a *AccountService) GetAllAccounts(ctx context.Context) ([]*entities.Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements AccountServiceI.
func (a *AccountService) UpdateAccount(ctx context.Context, in UpdateAccountIn) error {
	panic("unimplemented")
}
