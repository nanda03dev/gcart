package models

import (
	"github.com/nanda03dev/go2ms/common"
)

type Order struct {
	DocId   string   `json:"docId" bson:"docId"`
	Amount  int      `json:"amount" bson:"amount"`
	UserID  string   `json:"userId" bson:"userID"`
	Code    string   `json:"code" bson:"code"`
	ItemIds []string `json:"itemIds" bson:"itemIds"`
}

var OrderGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "orders",
	IndexKeys:      []string{"userId"},
}
