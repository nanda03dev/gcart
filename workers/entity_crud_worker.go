package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
)

var chanCRUD = make(chan common.EventType)

func AddToChanCRUD(event common.EventType) {
	chanCRUD <- event
}

func StartCRUDWorker() {
	for {
		crudEvent := <-chanCRUD

		event := models.Event{
			EntityId:      crudEvent.EntityId,
			EntityType:    crudEvent.EntityType,
			OperationType: crudEvent.OperationType,
			ExpireTime:    time.Now().Add(1 * time.Minute),
		}

		repositories.AppRepositories.Event.Create(context.TODO(), event)
	}
}
