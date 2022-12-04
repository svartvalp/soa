package server

import (
	"github.com/gin-gonic/gin"
)

type (
	Server interface {
		Run() error
	}

	controller interface {
		SetV1Handlers(group *gin.RouterGroup)
	}
)
