package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	city.ID = primitive.NewObjectID()
	if err := c.cityService.CreateCity(city); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, city)
}

func (c *CityController) GetAllCities(ctx *gin.Context) {
	citys, err := c.cityService.GetAllCities()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, citys)
}

func (c *CityController) GetCityByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	city, err := c.cityService.GetCityByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, city)
}

func (c *CityController) UpdateCity(ctx *gin.Context) {
	var city models.City
	if err := ctx.ShouldBindJSON(&city); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idParam := ctx.Param("id")
	city.ID, _ = primitive.ObjectIDFromHex(idParam)
	if err := c.cityService.UpdateCity(city); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, city)
}

func (c *CityController) DeleteCity(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.cityService.DeleteCity(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
