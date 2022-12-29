package server

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/soa/service/logger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	gin         *gin.Engine
	address     string
	controllers []controller
}

func NewServer(
	address string,
	controllers ...controller,
) Server {
	g := gin.New()
	g.Use(func(context *gin.Context) {
		context = logger.GinContextWithLogger(context)
	})
	g.Use(After)
	srv := &server{
		gin:         gin.Default(),
		address:     address,
		controllers: controllers,
	}

	srv.setHandlers()

	return srv
}

func After(ctx *gin.Context) {
	log := logger.LoggerFromGinContext(ctx)
	errs := ctx.Errors

	if len(errs.Errors()) > 0 {
		log.Error(errs.Last().Error())
		return
	}
	log.Infof("finish success")
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
