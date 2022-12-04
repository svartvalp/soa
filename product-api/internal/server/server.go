package server

import (
	_ "github.com/soa/product-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/soa/product-api/internal/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	gin         *gin.Engine
	address     string
	controllers []controller
}

func NewServer(
	cfg *config.Config,
	controllers ...controller,
) Server {
	srv := &server{
		gin:         gin.Default(),
		address:     cfg.Server.Address,
		controllers: controllers,
	}

	srv.setHandlers()

	return srv
}

// @title   Product API
// @version 1.0

// @host     localhost:7002
// @BasePath /api/v1
func (s *server) setHandlers() {
	v1 := s.gin.Group("/api/v1")
	{
		for _, c := range s.controllers {
			c.SetV1Handlers(v1)
		}
	}

	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *server) Run() error {
	return s.gin.Run(s.address)
}
