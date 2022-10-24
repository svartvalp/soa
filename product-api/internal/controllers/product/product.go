package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	productService productService
}

func NewController(productService productService) *Controller {
	return &Controller{productService: productService}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {
	g := group.Group("/product")
	{
		g.GET("/list", c.list)
		g.POST("/", c.create)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200 {array} models.Product
// @Router  /product/list [get]
func (c *Controller) list(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "List")
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200 {integer} string "id"
// @Router  /product [post]
func (c *Controller) create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}
