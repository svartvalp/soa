package server

import (
	"runtime"

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
	g.Use(LogAfter)
	srv := &server{
		gin:         g,
		address:     address,
		controllers: controllers,
	}

	srv.setHandlers()

	return srv
}

func LogBefore(ctx *gin.Context) {
	pc, _, _, _ := runtime.Caller(3)
	name := runtime.FuncForPC(pc).Name()
	log := logger.LoggerFromGinContext(ctx)
	log.Infof("%s start", name)
}

func LogAfter(ctx *gin.Context) {
	ctx.Next()
	pc, _, _, _ := runtime.Caller(3)
	name := runtime.FuncForPC(pc).Name()
	log := logger.LoggerFromGinContext(ctx)
	errs := ctx.Errors

	if len(errs.Errors()) > 0 {
		log.Errorf("%s finish with err: %v", name, errs.Last().Error())
		return
	}
	log.Infof("%s finish success", name)
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
