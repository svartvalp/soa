package indexer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	indexerService indexerService
}

func New(productService indexerService) *Controller {
	return &Controller{indexerService: productService}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {
	g := group.Group("/indexer")
	{
		g.GET("/regenerate", c.regenerate)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200
// @Router  /indexer/regenerate [get]
func (c *Controller) regenerate(ctx *gin.Context) {
	err := c.indexerService.Regenerate(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
