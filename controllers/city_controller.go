package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/gcart/common"
	"github.com/nanda03dev/gcart/global_constant"
	"github.com/nanda03dev/gcart/models"
	"github.com/nanda03dev/gcart/services"
)

type CityController struct {
	cityService services.CityService
}

func NewCityController(cityService services.CityService) *CityController {
	return &CityController{cityService}
}

func (c *CityController) CreateCity(ctx *gin.Context) {
	var city models.City
	if err := ctx.ShouldBindJSON(&city); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	city, err := c.cityService.CreateCity(city)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.ENTITY_CREATED_SUCCESSFULLY, city.DocId))
}

func (c *CityController) GetAllCities(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	cities, err := c.cityService.GetAllCities(requestFilterBody)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, cities))
}

func (c *CityController) GetCityByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	city, err := c.cityService.GetCityByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, city))
}

func (c *CityController) UpdateCity(ctx *gin.Context) {
	var city models.City

	if err := ctx.ShouldBindJSON(&city); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	fmt.Printf("\n controller city %v ", city)

	city.DocId = ctx.Param("id")
	if err := c.cityService.UpdateCity(city); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_UPDATED_SUCCESSFULLY, nil))
}

func (c *CityController) DeleteCity(ctx *gin.Context) {
	DocId := ctx.Param("id")

	if err := c.cityService.DeleteCity(DocId); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_DELETED_SUCCESSFULLY, nil))
}
