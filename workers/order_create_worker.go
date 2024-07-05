package workers

import "go.mongodb.org/mongo-driver/bson/primitive"

var channel = make(chan primitive.ObjectID)
func AddToOrderWorkerChan(orderId primitive.ObjectID) {
	channel <- orderId
}

func StartOrderWorker() {
	for {
		v, ok := <-channel
		if !ok {
			break
		}
		println("created order id  ", v.String())
	}
}
