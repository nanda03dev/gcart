package models

import (
	"github.com/nanda03dev/go2ms/common"
)

type Item struct {
	DocId  string `json:"docId" bson:"docId"`
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
	Status bool   `bson:"status"`
}

var ItemGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "items",
	IndexKeys:      []string{},
}
