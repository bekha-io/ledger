package v1

import (
	"github.com/bekha-io/ledger/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	Svc    services.ServiceFacade
	Logger zerolog.Logger
}

func (s *Server) Register(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("", s.GetAllAccounts)
			accounts.POST("", s.CreateAccount)
		}

		entries := v1.Group("/entries")
		{
			entries.GET("/:id", s.GetEntryByID)
			entries.POST("", s.CreateEntry)
		}
	}
}

func NewServer(svc services.ServiceFacade) *Server {
	return &Server{
		Svc:    svc,
		Logger: zerolog.Logger{}.With().Caller().Str("component", "v1.Server").Logger(),
	}
}
