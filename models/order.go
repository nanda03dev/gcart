package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type Order struct {
	DocId   string            `json:"docId" bson:"docId"`
	Amount  int               `json:"amount" bson:"amount"`
	UserID  string            `json:"userId" bson:"userID"`
	Code    common.StatusCode `json:"code" bson:"code"`
	ItemIds []string          `json:"itemIds" bson:"itemIds"`
}

var OrderGnosql = gnosql_client.CollectionInput{
	CollectionName: "orders",
	IndexKeys:      []string{"userId"},
}

func (order Order) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":   order.DocId,
		"amount":  order.Amount,
		"userId":  order.UserID,
		"code":    order.Code,
		"itemIds": order.ItemIds,
	}
}

func (order Order) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      order.DocId,
		EntityType:    global_constant.ENTITY_ORDER,
		OperationType: operationType,
	}
}
