package characteristic

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	characteristicService characteristicService
}

func NewController(characteristicService characteristicService) *Controller {
	return &Controller{characteristicService: characteristicService}
}

func (c *Controller) SetV1Handlers(group *gin.RouterGroup) {

}
