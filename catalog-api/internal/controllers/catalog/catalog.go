package catalog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soa/catalog-api/internal/models"
)

type Controller struct {
	catalogService catalogService
}

func New(catalogService catalogService) *Controller {
	return &Controller{
		catalogService: catalogService,
	}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {
	g := group.Group("/catalog")
	{
		g.GET("/brand/list", c.brandList)
		g.GET("/category/list", c.categoryList)
		g.POST("/product/list", c.productList)
		g.GET("/image/:image", c.getImage)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200 {array} string
// @Router  /catalog/brand/list [get]
func (c *Controller) brandList(ctx *gin.Context) {
	brands, err := c.catalogService.BrandList(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, brands)
}

// @Schemes
// @Accept  json
// @Param   filter body models.Filter true "List product filters"
// @Produce json
// @Success 200 {array} models.ProductInfo
// @Router  /catalog/product/list [post]
func (c *Controller) productList(ctx *gin.Context) {
	var filter models.Filter
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	products, err := c.catalogService.GetProducts(ctx, &filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200 {array} models.Category
// @Router  /catalog/category/list [get]
func (c *Controller) categoryList(ctx *gin.Context) {
	cats, err := c.catalogService.CategoryList(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, cats)
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200
// @Router  /catalog/image [get]
func (c *Controller) getImage(ctx *gin.Context) {
	name := ctx.Param("image")
	if name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "empty image name")
	}
	img, err := c.catalogService.GetImage(ctx, name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Data(200, img.ContentType, img.Data)
}
