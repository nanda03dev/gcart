package workers

import (
	"context"
	"time"

	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var channel = make(chan primitive.ObjectID)

func AddTo3201Chan(orderId primitive.ObjectID) {
	channel <- orderId
}

func Start3201Worker() {
	for {
		v := <-channel
		event := models.Event{
			ID:         primitive.NewObjectID(),
			OrderId:    v.String(),
			ExpireTime: time.Now().Add(1 * time.Minute),
		}
		repositories.AppRepositories.Event.Create(context.TODO(), event)
	}
}
