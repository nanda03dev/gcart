package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/services"
)

type RefundPaymentController struct {
	refundPaymentService services.RefundPaymentService
}

func NewRefundPaymentController(refundPaymentService services.RefundPaymentService) *RefundPaymentController {
	return &RefundPaymentController{refundPaymentService}
}

func (c *RefundPaymentController) CreateRefundPayment(ctx *gin.Context) {
	var refundPayment models.RefundPayment
	if err := ctx.ShouldBindJSON(&refundPayment); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	refundPayment, err := c.refundPaymentService.CreateRefundPayment(refundPayment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.ENTITY_CREATED_SUCCESSFULLY, refundPayment.DocId))
}

func (c *RefundPaymentController) GetAllRefundPayments(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	refundPayments, err := c.refundPaymentService.GetAllRefundPayments(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, refundPayments))

}

func (c *RefundPaymentController) GetRefundPaymentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	refundPayment, err := c.refundPaymentService.GetRefundPaymentByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_FETCHED_SUCCESSFULLY, refundPayment))
}

func (c *RefundPaymentController) UpdateRefundPayment(ctx *gin.Context) {
	var refundPayment models.RefundPayment
	if err := ctx.ShouldBindJSON(&refundPayment); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	refundPayment.DocId = ctx.Param("id")

	if err := c.refundPaymentService.UpdateRefundPayment(refundPayment); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_UPDATED_SUCCESSFULLY, nil))
}

func (c *RefundPaymentController) DeleteRefundPayment(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.refundPaymentService.DeleteRefundPayment(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_DELETED_SUCCESSFULLY, nil))
}

func (c *RefundPaymentController) ConfirmRefundPayment(ctx *gin.Context) {
	var refundPaymentConfirmBody common.RefundPaymentConfirmBody

	if err := ctx.ShouldBindJSON(&refundPaymentConfirmBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	if err := c.refundPaymentService.ConfirmRefundPayment(refundPaymentConfirmBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.ENTITY_CONFIRMED_SUCCESSFULLY, nil))
}
