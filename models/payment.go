package models

import "github.com/nanda03dev/gnosql_client"

type Payment struct {
	DocId   string `json:"docId" bson:"docId"`
	OrderId string `json:"orderId" bson:"orderId"`
	Name    string `json:"name" bson:"name"`
	Amount  int    `json:"amount" bson:"amount"`
	Code    string `json:"code" bson:"code"`
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
