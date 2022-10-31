package characteristic

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soa/product-api/internal/controllers/dto"
	"github.com/soa/product-api/internal/models"
)

type Controller struct {
	characteristicService characteristicService
}

func NewController(characteristicService characteristicService) *Controller {
	return &Controller{characteristicService: characteristicService}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {
	g := group.Group("/characteristic")
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
// @Success 200 {array} models.Characteristic
// @Router  /characteristic/list [get]
func (c *Controller) list(ctx *gin.Context) {
	res, err := c.characteristicService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}

// @Schemes
// @Accept  json
// @Param   products body dto.CreateCharacteristicReq true "Create characteristic"
// @Produce json
// @Success 200
// @Router  /characteristic/create [post]
func (c *Controller) create(ctx *gin.Context) {
	var characteristic dto.CreateCharacteristicReq

	err := ctx.ShouldBindJSON(&characteristic)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	id, err := c.characteristicService.Create(ctx, &characteristic)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, id)
}

// @Schemes
// @Accept  json
// @Param   product body models.Characteristic true "Update characteristic by id"
// @Produce json
// @Success 200
// @Router  /characteristic/update [put]
func (c *Controller) update(ctx *gin.Context) {
	var characteristic *models.Characteristic

	err := ctx.ShouldBindJSON(characteristic)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	err = c.characteristicService.Update(ctx, characteristic)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, err)
}

// @Schemes
// @Accept  json
// @Param   id query integer true "Delete characteristic by id"
// @Produce json
// @Success 200
// @Router  /characteristic/{id}/delete [delete]
func (c *Controller) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err = c.characteristicService.Delete(ctx, int64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, id)
}
