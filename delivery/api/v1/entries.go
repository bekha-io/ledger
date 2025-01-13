package v1

import (
	"github.com/bekha-io/ledger/domain/entities"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetEntryByID(c *gin.Context) {
}

func (s *Server) CreateEntry(c *gin.Context) {
	var in = struct {
		Transactions []struct {
			entities.Transaction
			ID uint `json:"-"`
		} `json:"transactions"`
	}{}

	if err := c.BindJSON(&in); err != nil {
		handleError(c, err)
		return
	}

	var transactions []*entities.Transaction
	for _, inTr := range in.Transactions {
		transactions = append(transactions, &entities.Transaction{
			AccountID:   inTr.AccountID,
			Description: inTr.Description,
			Type:        inTr.Type,
			Amount:      inTr.Amount,
		})
	}

	entry, err := s.Svc.Entries.CreateEntry(c, transactions)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, entry)
}
