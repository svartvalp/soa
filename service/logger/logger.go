package logger

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	k     string
	sugar *zap.SugaredLogger
)

type key struct {
}

func init() {
	b, _ := json.Marshal(key{})
	k = string(b)

	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()
}

func GinContextWithLogger(ctx *gin.Context) *gin.Context {
	ctx.Set(k, sugar)
	return ctx
}

func LoggerFromGinContext(ctx *gin.Context) *zap.SugaredLogger {
	if a, ok := ctx.Get(k); ok {
		if l, ok := a.(*zap.SugaredLogger); ok {
			return l
		}
	}
	return sugar
}

func ContextWithLogger(ctx context.Context) context.Context {
	return context.WithValue(ctx, k, sugar)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	if a := ctx.Value(key{}); a != nil {
		if l, ok := a.(*zap.SugaredLogger); ok {
			return l
		}
	}
	return sugar
}
