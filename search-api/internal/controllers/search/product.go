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
	products = append(products, models.ProductInfo{
		ID:              1,
		CategoryID:      1,
		Price:           100,
		Name:            "Погружной блендер REDMOND RHB-2941, белый",
		Description:     "Блендер Redmond RHB-2941 - новая компактная и многофункциональная модель \"3 в 1\" привлекательного белого  цвета из экологичной серии, которая помимо блендера является и измельчителем, и миксером. Насладитесь  мощностью, производительностью и почти бесшумной работой этого современного кухонного прибора! \n\n Блендер RHB-2941 оснащён инновационным сверхострым ножом S-образной формы и лезвиями из  высококачественной нержавеющей стали. Это устройство измельчит и равномерно смешает ингредиенты для  детского питания, освежающих коктейлей, супов-пюре, домашних соусов, жидкого теста, кремов, муссов,  изысканных десертов, очень быстро раздробит сухие продукты и смеси, измельчит мясо, твёрдые сыры, овощи,  травы, чеснок, орехи, способно моментально взбить сливки и яйца. \n",
		Brand:           "REDMOND",
		Image:           "1",
		Characteristics: nil,
		Categorys: []models.Category{
			{
				ID:          5,
				Name:        "Погружные",
				Description: "",
				ParentID:    4,
				Level:       3,
			},
			{
				ID:          4,
				Name:        "Блендеры",
				Description: "",
				ParentID:    3,
				Level:       2,
			},
			{
				ID:          3,
				Name:        "Бытовая техника",
				Description: "",
				ParentID:    0,
				Level:       1,
			},
		},
	})
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
