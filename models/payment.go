package models

import (
	"github.com/nanda03dev/go2ms/common"
)

type Payment struct {
	DocId   string `json:"docId" bson:"docId"`
	OrderId string `json:"orderId" bson:"orderId"`
	Name    string `json:"name" bson:"name"`
	Amount  int    `json:"amount" bson:"amount"`
	Code    string `json:"code" bson:"code"`
}

var PaymentGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "payments",
	IndexKeys:      []string{"orderId"},
}
