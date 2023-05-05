package handler

import (
	"challange10-dts/dto"
	"challange10-dts/entity"
	"challange10-dts/pkg/errs"
	"challange10-dts/pkg/helpers"
	"challange10-dts/service"
	"net/http"

	_ "challange10-dts/entity"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

// CreateNewProduct godoc
// @Tags products
// @Description Create New Product Data
// @ID create-new-product
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewProductRequest true "request body json"
// @Success 201 {object} dto.NewProductRequest
// @Router /products [post]
func (m productHandler) CreateProduct(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newMovie, err := m.productService.CreateProduct(user.Id, productRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newMovie)
}

func (m productHandler) UpdateProductById(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	productId, err := helpers.GetParamId(c, "productId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		// log.Println(err)
		return
	}

	response, err := m.productService.UpdateProductById(productId, productRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		// log.Println(err)
		return
	}

	c.JSON(response.StatusCode, response)
}
