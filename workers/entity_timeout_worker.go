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

			if event.EntityType == global_constant.ENTITY_ORDER {
				order, _ := orderRepository.GetByID(context.TODO(), event.EntityId)

				if order.Code == global_constant.ORDER_INITIATED {
					order.Code = global_constant.ORDER_TIMEOUT
					orderRepository.Update(context.TODO(), order.DocId, order)
					eventRepository.Delete(context.TODO(), event.DocId)
				}
			}
			if event.EntityType == global_constant.ENTITY_PAYMENT {
				payment, _ := paymentRepository.GetByID(context.TODO(), event.EntityId)

				if payment.Code == global_constant.PAYMENT_INITIATED {
					payment.Code = global_constant.PAYMENT_TIMEOUT
					paymentRepository.Update(context.TODO(), payment.DocId, payment)
					eventRepository.Delete(context.TODO(), event.DocId)
				}
			}

		}

		time.Sleep(time.Second * 10)
	}
}
