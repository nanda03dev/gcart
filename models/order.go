package models

import "github.com/nanda03dev/gnosql_client"

type Order struct {
	DocId   string   `json:"docId" bson:"docId"`
	Amount  int      `json:"amount" bson:"amount"`
	UserID  string   `json:"userId" bson:"userID"`
	Code    string   `json:"code" bson:"code"`
	ItemIds []string `json:"itemIds" bson:"itemIds"`
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
