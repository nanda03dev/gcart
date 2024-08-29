package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/repositories"
	"github.com/nanda03dev/gcart/src/services"
)

func StartEntityTimeoutWorker() {
	var eventRepository = repositories.AppRepositories.Event

	var orderService = services.AppServices.Order
	var itemService = services.AppServices.Item
	var paymentService = services.AppServices.Payment
	var eventService = services.AppServices.Event

	for {
		eventLimit := 100

		timeoutFilter := common.FiltersBodyType{
			{Key: "checkProcess", Value: global_constant.CHECK_TIMEOUT},
		}

		events, _ := eventRepository.GetAll(context.TODO(), timeoutFilter, common.SortBodyType{Key: "createdAt", Order: 1}, eventLimit)

		for _, event := range events {

			if !models.IsEventTimeExpired(event.EntityType, event.CreatedAt) {
				continue
			}

			if event.EntityType == global_constant.ENTITY_ORDER {
				isOrderTimedout := orderService.UpdateOrderTimeout(event.EntityId)
				if isOrderTimedout {
					orderIdFilter := common.FiltersBodyType{
						{Key: "orderId", Value: event.EntityId},
					}
					itemService.UpdateItemsTimeout(orderIdFilter)
				}
			}

			if event.EntityType == global_constant.ENTITY_PAYMENT {
				paymentService.UpdatePaymentTimeout(event.EntityId)
			}

			event.CheckProcess = global_constant.CHECK_TIMEOUT_DONE
			eventService.UpdateEvent(event)
		}

		time.Sleep(time.Second * 10)
	}
}
