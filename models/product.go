package models

import "github.com/nanda03dev/go2ms/common"

type Product struct {
	DocId  string `json:"docId" bson:"docId"`
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
	Status bool   `bson:"status"`
}

var ProductGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "products",
	IndexKeys:      []string{},
}
