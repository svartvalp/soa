package server

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/svartvalp/soa/service/logger"
	"github.com/svartvalp/soa/service/server/internal/metrics"
)

func WithLogger() func(g *gin.Engine) {
	return func(g *gin.Engine) {
		g.Use(func(context *gin.Context) {
			context = logger.GinContextWithLogger(context)
		})
		g.Use(func(ctx *gin.Context) {
			if notApi(ctx) {
				return
			}
			log := logger.LoggerFromGinContext(ctx)
			log.Infof("%s start", ctx.HandlerName())
			ctx.Next()
			errs := ctx.Errors
			if len(errs.Errors()) > 0 {
				log.Errorf("%s finish with err: %v", ctx.HandlerName(), errs.Last().Error())
				return
			}
			log.Infof("%s finish success", ctx.HandlerName())
		})
	}
}

func WithMetrics() func(g *gin.Engine) {
	return func(g *gin.Engine) {
		g.GET("/metrics", func(ctx *gin.Context) {
			h := promhttp.Handler()
			h.ServeHTTP(ctx.Writer, ctx.Request)
		})
		g.Use(func(ctx *gin.Context) {
			if notApi(ctx) {
				return
			}
			now := time.Now()
			ctx.Next()
			duration := time.Since(now)
			metrics.ResponseTimeHistogram.WithLabelValues(
				ctx.FullPath(),
				ctx.Request.Method,
				strconv.Itoa(ctx.Writer.Status()),
			).Observe(duration.Seconds())
		})
	}
}

func notApi(ctx *gin.Context) bool {
	return strings.Contains(ctx.FullPath(), "metrics") || strings.Contains(ctx.FullPath(), "docs")
}
