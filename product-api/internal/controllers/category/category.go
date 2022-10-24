package category

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	categoryService categoryService
}

func NewController(categoryService categoryService) *Controller {
	return &Controller{categoryService: categoryService}
}

func (c Controller) SetV1Handlers(group *gin.RouterGroup) {
	// g := group.Group("/category")
	// {
	//	g.GET("")
	// }
}
