package consumers

import (
	"fmt"

	"github.com/nanda03dev/gcart/config"
	"github.com/nanda03dev/gcart/message"
	"github.com/nanda03dev/gque_client"
)

func StartAllCrudGqueConsumer() {
	consumer := make(chan gque_client.MessageType, 10000)

	consumerRequestType := gque_client.ConsumerRequestType{
		QueueName: message.CrudEventGque.Name,
	}

	err := config.GqueClient.Consume(consumerRequestType, consumer)

	if err != nil {
		fmt.Printf("\n Gque consumer connection failed for gque: %v ", message.CrudEventGque.Name)
		return
	} else {
		fmt.Printf("\n Gque consumer started successfully for gque: %v ", message.CrudEventGque.Name)
	}

	for {
		message := <-consumer
		fmt.Printf("\n Received message: %v ", message)
	}
}
