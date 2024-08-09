package message

import (
	"fmt"

	"github.com/nanda03dev/gcart/config"
	"github.com/nanda03dev/gque_client"
)

var CrudEventGque = gque_client.Queue{
	Name: "all-crud-events",
	Time: 20,
}

var QueueList = []gque_client.Queue{
	CrudEventGque,
}

func InitializeGque() {
	for _, queue := range QueueList {
		result, err := config.GqueClient.CreateQueue(queue)
		if err == nil {
			fmt.Printf("\n Queue %v successfully created ", result)
		} else {
			fmt.Printf("\n Queue already exists \n")
		}
	}
}

func PushMessageToGque(gque gque_client.Queue, message gque_client.MessageType) {
	config.GqueClient.PushMessage(gque_client.QueueMessageType{
		Name: gque.Name,
		Data: message,
	})
}
