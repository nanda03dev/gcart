package workers

import (
	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/services"
)

func StartPaymentRefundWorker() {
	for {
		crudEvent := <-common.ChannelPaymentRefund

		var orderService = services.AppServices.Order
		var paymentService = services.AppServices.Payment
		var refundPaymentService = services.AppServices.RefundPayment

		order, _ := orderService.GetOrderByID(crudEvent.EntityId)

		if order.StatusCode != global_constant.ORDER_CONFIRMED && order.StatusCode != global_constant.ORDER_TIMEOUT {
			return
		}

		orderPayments, _ := paymentService.GetAllPaymentsByOrderId(order.DocId)

		var orderAmount = order.Amount
		var paymentTotalAmount int

		for _, payment := range orderPayments {
			if payment.StatusCode == global_constant.PAYMENT_CONFIRMED {
				paymentTotalAmount = paymentTotalAmount + payment.Amount
			}
		}
		var refundAmount = 0

		if order.StatusCode == global_constant.ORDER_CONFIRMED {
			refundAmount = paymentTotalAmount - orderAmount
		} else {
			refundAmount = paymentTotalAmount
		}

		if refundAmount < 1 {
			return
		}

		newRefundPayment := models.RefundPayment{
			OrderId: order.DocId,
			Amount:  refundAmount,
		}

		refundPaymentService.CreateRefundPayment(newRefundPayment)

	}
}
