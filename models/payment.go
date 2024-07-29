package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type Payment struct {
	DocId   string            `json:"docId" bson:"docId"`
	OrderId string            `json:"orderId" bson:"orderId"`
	Name    string            `json:"name" bson:"name"`
	Amount  int               `json:"amount" bson:"amount"`
	Code    common.StatusCode `json:"code" bson:"code"`
}

var PaymentGnosql = gnosql_client.CollectionInput{
	CollectionName: "payments",
	IndexKeys:      []string{"orderId"},
}

func (payment Payment) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":   payment.DocId,
		"orderId": payment.OrderId,
		"name":    payment.Name,
		"amount":  payment.Amount,
		"code":    payment.Code,
	}
}

func (payment Payment) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      payment.DocId,
		EntityType:    global_constant.ENTITY_PAYMENT,
		OperationType: operationType,
	}
}
