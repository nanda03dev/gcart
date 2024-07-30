package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/services"
	"github.com/nanda03dev/go2ms/utils"
)

func StartEntityTimeoutWorker() {
	var eventRepository = repositories.AppRepositories.Event

	var orderService = services.AppServices.Order
	var itemService = services.AppServices.Item
	var paymentService = services.AppServices.Payment
	var eventService = services.AppServices.Event

	for {
		eventLimit := 100

		events, _ := eventRepository.GetAll(context.TODO(), nil, common.SortBodyType{Key: "createdAt", Order: 1}, eventLimit)

		for _, event := range events {

			if !utils.IsEventTimeExpired(event.EntityType, event.CreatedAt) {
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

			eventService.DeleteEvent(event.DocId)
		}

		time.Sleep(time.Second * 10)
	}
}
