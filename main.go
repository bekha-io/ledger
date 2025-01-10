package main

import (
	v1 "github.com/bekha-io/ledger/delivery/api/v1"
	"github.com/bekha-io/ledger/domain/repository"
	"github.com/bekha-io/ledger/domain/services"
	"github.com/bekha-io/ledger/infrastructure/repository/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creating repositories
	accountsRepo := &memory.MemoryAccountRepository{}
	entriesRepo := &memory.MemoryEntryRepository{}
	repo := repository.RepoFacade{
		Accounts: accountsRepo,
		Entries:  entriesRepo,
	}

	// Creating services
	svc := services.ServiceFacade{
		Accounts: &services.AccountService{Repo: repo},
		Entries:  &services.EntryService{Repo: repo},
	}

	// Creating gin engine
	e := gin.New()
	e.Use(gin.Recovery())

	// Registering all the endpoints
	server := v1.NewServer(svc)
	server.Register(e)

	// Running the application
	e.Run(":8080")
}
