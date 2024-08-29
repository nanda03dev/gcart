package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/services"
)

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) *PaymentController {
	return &PaymentController{paymentService}
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	payment, err := c.paymentService.CreatePayment(payment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.ENTITY_CREATED_SUCCESSFULLY, payment.DocId))
}

func (c *PaymentController) GetAllPayments(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	payments, err := c.paymentService.GetAllPayments(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, payments))

}

func (c *PaymentController) GetPaymentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	payment, err := c.paymentService.GetPaymentByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, payment))
}

func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	payment.DocId = ctx.Param("id")

	if err := c.paymentService.UpdatePayment(payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_UPDATED_SUCCESSFULLY, nil))
}

func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.paymentService.DeletePayment(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_DELETED_SUCCESSFULLY, nil))
}

func (c *PaymentController) ConfirmPayment(ctx *gin.Context) {
	var paymentConfirmBody common.PaymentConfirmBody

	if err := ctx.ShouldBindJSON(&paymentConfirmBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	if err := c.paymentService.ConfirmPayment(paymentConfirmBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_CONFIRMED_SUCCESSFULLY, nil))
}
