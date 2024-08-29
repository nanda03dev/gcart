package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/services"
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
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	order, err := c.orderService.CreateOrder(order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.ENTITY_CREATED_SUCCESSFULLY, order.DocId))
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	orders, err := c.orderService.GetAllOrders(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, orders))

}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	order, err := c.orderService.GetOrderByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, order))
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	order.DocId = ctx.Param("id")

	if err := c.orderService.UpdateOrder(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_UPDATED_SUCCESSFULLY, nil))
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.orderService.DeleteOrder(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_DELETED_SUCCESSFULLY, nil))
}

func (c *OrderController) ConfirmOrder(ctx *gin.Context) {
	var orderConfirmBody common.OrderConfirmBody

	if err := ctx.ShouldBindJSON(&orderConfirmBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	if err := c.orderService.ConfirmOrder(orderConfirmBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_CONFIRMED_SUCCESSFULLY, nil))
}
