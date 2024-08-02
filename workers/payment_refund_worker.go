package workers

import (
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
)

func StartPaymentRefundWorker() {
	for {
		crudEvent := <-common.ChannelPaymentRefund

		var orderService = services.AppServices.Order
		var paymentService = services.AppServices.Payment
		var refundPaymentService = services.AppServices.RefundPayment

		order, _ := orderService.GetOrderByID(crudEvent.EntityId)

		if order.StatusCode != global_constant.ORDER_CONFIRMED {
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
		refundAmount := paymentTotalAmount - orderAmount

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
