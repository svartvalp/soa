package server

import (
	"github.com/gin-gonic/gin"
)

type (
	Server interface {
		Run() error
	}

	Controller interface {
		SetV1Handlers(group *gin.RouterGroup)
	}
)
