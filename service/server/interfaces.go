package server

import (
	"github.com/gin-gonic/gin"
)

type (
	Server interface {
		Run() error
	}

	option func(engine *gin.Engine)

	Controller interface {
		SetV1Handlers(group *gin.RouterGroup)
	}
)
