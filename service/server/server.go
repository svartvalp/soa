package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	server struct {
		gin         *gin.Engine
		address     string
		controllers []controller
	}

	Config struct {
		Host        string
		Port        int
		controllers []controller
	}
)

func NewServer(
	cfg *Config,
	options ...option,
) Server {
	g := gin.New()

	for _, o := range options {
		o(g)
	}

	srv := &server{
		gin:         g,
		address:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		controllers: cfg.controllers,
	}

	srv.setHandlers()

	return srv
}

func (s *server) setHandlers() {
	v1 := s.gin.Group("/api/v1")
	{
		for _, c := range s.controllers {
			c.SetV1Handlers(v1)
		}
	}

	s.gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *server) Run() error {
	return s.gin.Run(s.address)
}
