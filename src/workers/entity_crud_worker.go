package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/config"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/message_queue"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/repositories"
	"github.com/nanda03dev/gcart/src/services"
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/gque_client"
)

func StartCRUDWorker() {
	for {
		crudEvent := <-common.ChannelCRUD

		var cityService = services.AppServices.City
		var userService = services.AppServices.User
		var productService = services.AppServices.Product
		var orderService = services.AppServices.Order
		var itemService = services.AppServices.Item
		var paymentService = services.AppServices.Payment
		var refundPaymentService = services.AppServices.RefundPayment

		event := models.Event{
			EntityId:      crudEvent.EntityId,
			EntityType:    crudEvent.EntityType,
			OperationType: crudEvent.OperationType,
			CreatedAt:     time.Now(),
			CheckProcess:  crudEvent.CheckProcess,
		}

		collectionName := models.GetGnosqlCollection(crudEvent.EntityType).CollectionName
		entityGnosql := config.GnoSQLDB.Collections[collectionName]
		var docmentToCreate gnosql_client.Document

		switch event.EntityType {
		case global_constant.ENTITY_CITY:
			{
				city, _ := cityService.GetCityByID(event.EntityId)
				docmentToCreate = city.ToDocument()
			}
		case global_constant.ENTITY_USER:
			{
				user, _ := userService.GetUserByID(event.EntityId)
				docmentToCreate = user.ToDocument()
			}
		case global_constant.ENTITY_PRODUCT:
			{
				product, _ := productService.GetProductByID(event.EntityId)
				docmentToCreate = product.ToDocument()
			}
		case global_constant.ENTITY_ORDER:
			{
				order, _ := orderService.GetOrderByID(event.EntityId)
				docmentToCreate = order.ToDocument()
			}
		case global_constant.ENTITY_ITEM:
			{
				item, _ := itemService.GetItemByID(event.EntityId)
				docmentToCreate = item.ToDocument()
			}
		case global_constant.ENTITY_PAYMENT:
			{
				payment, _ := paymentService.GetPaymentByID(event.EntityId)
				docmentToCreate = payment.ToDocument()
			}
		case global_constant.ENTITY_REFUND_PAYMENT:
			{
				refundPayment, _ := refundPaymentService.GetRefundPaymentByID(event.EntityId)
				docmentToCreate = refundPayment.ToDocument()
			}
		}

		switch event.OperationType {

		case global_constant.OPERATION_CREATE:
			{
				entityGnosql.Create(docmentToCreate)
			}
		case global_constant.OPERATION_UPDATE, global_constant.OPERATION_CONFIRMED:
			{
				entityGnosql.Update(event.EntityId, docmentToCreate)
			}
		case global_constant.OPERATION_DELETE:
			{
				entityGnosql.Delete(event.EntityId)
			}
		}

		if models.IsRequireToStoreEvent(event.EntityType) {
			repositories.AppRepositories.Event.Create(context.TODO(), event)
		}

		queueMessage := gque_client.MessageType{
			"data": crudEvent,
		}
		message_queue.PushMessageToGque(message_queue.CrudEventGque, queueMessage)
	}
}
