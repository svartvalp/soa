package category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type Controller struct {
	categoryService categoryService
}

func NewController(categoryService categoryService) *Controller {
	return &Controller{categoryService: categoryService}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {
	g := group.Group("/ategory")
	{
		g.GET("/list", c.list)
		g.POST("/create", c.create)
		g.PUT("/update", c.update)
		g.DELETE("/:id/delete", c.delete)
	}
}

// @Schemes
// @Accept  json
// @Produce json
// @Success 200 {array} models.Category
// @Router  /category/list [get]
func (c *Controller) list(ctx *gin.Context) {
	res, err := c.categoryService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Schemes
// @Accept  json
// @Param   products body dto.CreateCategoryReq true "Create category"
// @Produce json
// @Success 200
// @Router  /category/create [post]
func (c *Controller) create(ctx *gin.Context) {
	var category dto.CreateCategoryReq

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	id, err := c.categoryService.Create(ctx, &category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, id)
}

// @Schemes
// @Accept  json
// @Param   product body models.Category true "Update category by id"
// @Produce json
// @Success 200
// @Router  /category/update [put]
func (c *Controller) update(ctx *gin.Context) {
	var category *models.Category

	err := ctx.ShouldBindJSON(category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err = c.categoryService.Update(ctx, category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, err)
}

// @Schemes
// @Accept  json
// @Param   id query integer true "Delete category by id"
// @Produce json
// @Success 200
// @Router  /category/{id}/delete [delete]
func (c *Controller) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err = c.categoryService.Delete(ctx, int64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, id)
}
