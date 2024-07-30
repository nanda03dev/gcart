package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	product, err := c.productService.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.DATA_CREATED_SUCCESSFULLY, product.DocId))
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	products, err := c.productService.GetAllProducts(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, products))
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	product, err := c.productService.GetProductByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, product))
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	product.DocId = ctx.Param("id")

	if err := c.productService.UpdateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_UPDATED_SUCCESSFULLY, nil))
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.productService.DeleteProduct(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_DELETED_SUCCESSFULLY, nil))
}
