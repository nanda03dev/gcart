package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type Order struct {
	DocId      string            `json:"docId" bson:"docId"`
	Amount     int               `json:"amount" bson:"amount"`
	UserID     string            `json:"userId" bson:"userID"`
	StatusCode common.StatusCode `json:"statusCode" bson:"statusCode"`
}

var OrderGnosql = gnosql_client.CollectionInput{
	CollectionName: "orders",
	IndexKeys:      []string{"userId"},
}

func (order Order) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      order.DocId,
		"amount":     order.Amount,
		"userId":     order.UserID,
		"statusCode": order.StatusCode,
	}
}

func (order Order) ToModel(orderDocument gnosql_client.Document) Order {
	return Order{
		DocId:      GetStringValue(orderDocument, "docId"),
		Amount:     GetIntegerValue(orderDocument, "amount"),
		UserID:     GetStringValue(orderDocument, "userId"),
		StatusCode: GetValue[common.StatusCode](orderDocument, "statusCode"),
	}
}

func (order Order) ToEvent(operationType common.OperationType) common.EventType {

	return common.EventType{
		EntityId:      order.DocId,
		EntityType:    global_constant.ENTITY_ORDER,
		OperationType: operationType,
		CheckProcess:  GetCheckProcess(global_constant.ENTITY_ORDER, operationType),
	}
}

func (order Order) ToUpdatedDocument(newOrder Order) Order {
	orderDocument := order.ToDocument()
	newOrderDocument := newOrder.ToDocument()

	// order statusCode should not be updated
	newOrderDocument["statusCode"] = orderDocument["statusCode"]

	for key, value := range newOrderDocument {
		if value != nil && value != "" {
			orderDocument[key] = value
		}
	}

	return order.ToModel(orderDocument)
}
