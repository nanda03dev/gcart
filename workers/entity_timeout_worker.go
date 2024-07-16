package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/repositories"
)

func StartEntityTimeoutWorker() {
	var orderRepository = repositories.AppRepositories.Order
	var paymentRepository = repositories.AppRepositories.Payment
	var eventRepository = repositories.AppRepositories.Event

	for {

		events, _ := eventRepository.GetAll(context.TODO(), nil, common.SortBodyType{Key: "ExpireTime", Order: 1}, nil)

		for _, event := range events {

			if time.Now().Before(event.ExpireTime) {
				continue
			}

			if event.EntityType == global_constant.Entities.Order {
				order, _ := orderRepository.GetByID(context.TODO(), event.EntityId)

				if order.Code == global_constant.OrderSuccessCode.ORDER_INITIATED {
					order.Code = global_constant.OrderErrorCode.ORDER_TIMEOUT
					orderRepository.Update(context.TODO(), order.ID, order)
					eventRepository.Delete(context.TODO(), event.ID)
				}
			}
			if event.EntityType == global_constant.Entities.Payment {
				payment, _ := paymentRepository.GetByID(context.TODO(), event.EntityId)

				if payment.Code == global_constant.PaymentSuccessCode.PAYMENT_INITIATED {
					payment.Code = global_constant.PaymentErrorCode.PAYMENT_TIMEOUT
					paymentRepository.Update(context.TODO(), payment.ID, payment)
					eventRepository.Delete(context.TODO(), event.ID)
				}
			}

		}

		time.Sleep(time.Second * 10)
	}
}
