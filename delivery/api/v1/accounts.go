package v1

import (
	"github.com/bekha-io/ledger/domain/services"
	"github.com/bekha-io/ledger/domain/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetAllAccounts(c *gin.Context) {
	accounts, err := s.Svc.Accounts.GetAllAccounts(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(200, accounts)
}

func (s *Server) CreateAccount(c *gin.Context) {
	var in = struct {
		ID          string            `json:"id" binding:"required"`
		Name        string            `json:"name" binding:"required"`
		Description string            `json:"description"`
		Type        types.AccountType `json:"type"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		handleError(c, err)
		return
	}

	_, err := s.Svc.Accounts.CreateAccount(c, services.CreateAccountIn{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Type:        in.Type,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Created successfully!"})
}
