package consumers

import (
	"fmt"

	"github.com/nanda03dev/gcart/src/config"
	"github.com/nanda03dev/gcart/src/message_queue"
	"github.com/nanda03dev/gque_client"
)

func StartAllCrudGqueConsumer() {
	consumer := make(chan gque_client.MessageType, 10000)

	consumerRequestType := gque_client.ConsumerRequestType{
		QueueName: message_queue.CrudEventGque.Name,
	}

	err := config.GqueClient.Consume(consumerRequestType, consumer)

	if err != nil {
		fmt.Printf("\n Gque consumer connection failed for gque: %v ", message_queue.CrudEventGque.Name)
		return
	} else {
		fmt.Printf("\n Gque consumer started successfully for gque: %v ", message_queue.CrudEventGque.Name)
	}

	for {
		message := <-consumer
		fmt.Printf("\n Received message: %v ", message)
	}
}
