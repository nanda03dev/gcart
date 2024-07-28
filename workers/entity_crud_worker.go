package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/config"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
)

var chanCRUD = make(chan common.EventType)

func AddToChanCRUD(event common.EventType) {
	chanCRUD <- event
}

func StartCRUDWorker() {
	for {
		crudEvent := <-chanCRUD
		var cityRepository = repositories.AppRepositories.City
		var userRepository = repositories.AppRepositories.User
		var productRepository = repositories.AppRepositories.Product

		var orderRepository = repositories.AppRepositories.Order
		var itemRepository = repositories.AppRepositories.Item
		var paymentRepository = repositories.AppRepositories.Payment

		event := models.Event{
			EntityId:      crudEvent.EntityId,
			EntityType:    crudEvent.EntityType,
			OperationType: crudEvent.OperationType,
			ExpireTime:    time.Now().Add(1 * time.Minute),
		}

		collectionName := utils.GetGnosqlCollection(crudEvent.EntityType).CollectionName
		entityGnosql := config.GnoSQLDB.Collections[collectionName]
		var docmentToCreate gnosql_client.Document

		switch event.EntityType {
		case global_constant.Entities.City:
			{
				city, _ := cityRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = city.ToDocument()
			}
		case global_constant.Entities.User:
			{
				user, _ := userRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = user.ToDocument()
			}
		case global_constant.Entities.Product:
			{
				product, _ := productRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = product.ToDocument()
			}
		case global_constant.Entities.Order:
			{
				order, _ := orderRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = order.ToDocument()
			}
		case global_constant.Entities.Item:
			{
				item, _ := itemRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = item.ToDocument()
			}
		case global_constant.Entities.Payment:
			{
				payment, _ := paymentRepository.GetByID(context.TODO(), event.EntityId)
				docmentToCreate = payment.ToDocument()
			}
		}

		switch event.OperationType {

		case global_constant.Operations.Create:
			{
				entityGnosql.Create(docmentToCreate)
			}
		case global_constant.Operations.Update:
			{
				entityGnosql.Update(event.EntityId, docmentToCreate)
			}
		case global_constant.Operations.Delete:
			{
				entityGnosql.Delete(event.EntityId)
			}
		}

		repositories.AppRepositories.Event.Create(context.TODO(), event)
	}
}
