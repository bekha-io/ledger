package services

import (
	"context"

	"github.com/bekha-io/ledger/domain/entities"
	"github.com/bekha-io/ledger/domain/errs"
	"github.com/bekha-io/ledger/domain/repository"
	"github.com/bekha-io/ledger/domain/types"
	"github.com/shopspring/decimal"
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

	// First we check if such account exists
	acc, err := a.Repo.Accounts.GetAccountByID(ctx, in.ID)
	if err == nil && acc != nil {
		return nil, errs.ErrAccountWithSuchIDExists
	}

	// Then we validate and save in repo
	account := &entities.Account{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Type:        in.Type,
		Balance:     decimal.Decimal{},
	}

	if err := account.Validate(); err != nil {
		return nil, err
	}

	if err := a.Repo.Accounts.SaveAccount(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}

// GetAccountByID implements AccountServiceI.
func (a *AccountService) GetAccountByID(ctx context.Context, id string) (*entities.Account, error) {
	panic("unimplemented")
}

// GetAllAccounts implements AccountServiceI.
func (a *AccountService) GetAllAccounts(ctx context.Context) ([]*entities.Account, error) {
	return a.Repo.Accounts.GetAllAccounts(ctx)
}

// UpdateAccount implements AccountServiceI.
func (a *AccountService) UpdateAccount(ctx context.Context, in UpdateAccountIn) error {
	panic("unimplemented")
}
