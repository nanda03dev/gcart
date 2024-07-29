package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type Product struct {
	DocId  string `json:"docId" bson:"docId"`
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
	Status bool   `bson:"status"`
}

var ProductGnosql = gnosql_client.CollectionInput{
	CollectionName: "products",
	IndexKeys:      []string{},
}

func (product Product) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":  product.DocId,
		"name":   product.Name,
		"amount": product.Amount,
		"status": product.Status,
	}
}

func (product Product) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      product.DocId,
		EntityType:    global_constant.ENTITY_PRODUCT,
		OperationType: operationType,
	}
}
