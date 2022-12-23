package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soa/search-api/internal/models"
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
		g.POST("/list", c.list)
		g.PUT("/update", c.regenerate)
		g.POST("/index", c.commitIndex)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Param   filter body models.Filter true "List product filters"
// @Success 200    {array} models.ProductInfo
// @Router  /product/list [post]
func (c *Controller) list(ctx *gin.Context) {
	var filter models.Filter
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ids, err := c.productService.List(ctx, &filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, ids)
}

// @Schemes
// @Accept  json
// @Param   products body  []models.ProductInfo true "Update products"
// @Produce json
// @Success 200
// @Router  /product/update [put]
func (c *Controller) regenerate(ctx *gin.Context) {
	var products []models.ProductInfo
	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err = c.productService.Regenerate(ctx, products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Status(http.StatusOK)
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200
// @Router  /product/index [post]
func (c *Controller) commitIndex(ctx *gin.Context) {
	err := c.productService.SwapIndex(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Status(http.StatusOK)
}
