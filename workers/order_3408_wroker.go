package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/repositories"
)

func Start3408Worker() {
	var orderRepository = repositories.AppRepositories.Order
	var eventRepository = repositories.AppRepositories.Event

	for {
		events, _ := eventRepository.GetAll(context.TODO(), nil, common.SortBody{Key: "ExpireTime", Order: 1}, nil)

		for _, event := range events {
			order, _ := orderRepository.GetByID(context.TODO(), event.OrderId)

			if order.Code == global_constant.ORDER_SUCCESS_STATUS_CODE.ORDER_INITIATED {
				if event.ExpireTime.Before(time.Now()) {
					orderRepository.Delete(context.TODO(), order.ID)
					eventRepository.Delete(context.TODO(), event.ID)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
