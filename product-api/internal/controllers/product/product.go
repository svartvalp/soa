package product

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
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
		g.POST("/create", c.create)
		g.PUT("/update", c.update)
		g.DELETE("/:id/delete", c.delete)
		g.POST("/add", c.loadImage)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Param   filter body    models.ProductFilters true "List product filters"
// @Success 200    {array} models.Product
// @Router  /product/list [post]
func (c *Controller) list(ctx *gin.Context) {
	req := &models.ProductFilters{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	res, err := c.productService.List(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}

// @Schemes
// @Accept  json
// @Param   products body dto.CreateProductReq true "Create product list"
// @Produce json
// @Success 200
// @Router  /product/create [post]
func (c *Controller) create(ctx *gin.Context) {
	var product dto.CreateProductReq

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	id, err := c.productService.Create(ctx, &product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, id)
}

// @Schemes
// @Accept  multipart/form-data
// @Param   file formData file true "Product image"
// @Produce json
// @Success 200
// @Router  /product/add [post]
func (c *Controller) loadImage(ctx *gin.Context) {
	formFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	f, err := formFile.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	b, err := io.ReadAll(f)

	err = c.productService.UI(ctx, &dto.Image{Body: b, Name: formFile.Filename})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Status(http.StatusOK)
}

// @Schemes
// @Accept  json
// @Param   product body models.Product true "Update product by id"
// @Produce jsons
// @Success 200
// @Router  /product/update [put]
func (c *Controller) update(ctx *gin.Context) {
	var products *models.Product

	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err = c.productService.Update(ctx, products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	ctx.Status(http.StatusOK)
}

// @Schemes
// @Accept  json
// @Param   id query integer true "Delete product by id"
// @Produce json
// @Success 200
// @Router  /product/{id}/delete [delete]
func (c *Controller) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err = c.productService.Delete(ctx, int64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	ctx.Status(http.StatusOK)
}
