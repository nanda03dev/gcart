package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.ID = primitive.NewObjectID()
	if err := c.orderService.CreateOrder(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	order, err := c.orderService.GetOrderByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idParam := ctx.Param("id")
	order.ID, _ = primitive.ObjectIDFromHex(idParam)
	if err := c.orderService.UpdateOrder(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.orderService.DeleteOrder(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
