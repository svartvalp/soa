package server

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/soa/service/logger"
)

func logBefore(ctx *gin.Context) {
	log := logger.LoggerFromGinContext(ctx)
	log.Infof("%s start", ctx.HandlerName())
}

func logAfter(ctx *gin.Context) {
	ctx.Next()
	log := logger.LoggerFromGinContext(ctx)
	errs := ctx.Errors
	if len(errs.Errors()) > 0 {
		log.Errorf("%s finish with err: %v", ctx.HandlerName(), errs.Last().Error())
		return
	}
	log.Infof("%s finish success", ctx.HandlerName())
}

func WithLogger(g *gin.Engine) {
	g.Use(func(context *gin.Context) {
		context = logger.GinContextWithLogger(context)
	})
	g.Use(logBefore)
	g.Use(logAfter)
}
