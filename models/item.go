package models

import (
	"github.com/nanda03dev/gnosql_client"
)

type Item struct {
	DocId  string `json:"docId" bson:"docId"`
	Name   string `json:"name" bson:"name"`
	Amount int    `json:"amount" bson:"amount"`
	Status bool   `bson:"status"`
}

var ItemGnosql = gnosql_client.CollectionInput{
	CollectionName: "items",
	IndexKeys:      []string{},
}

func (item Item) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":  item.DocId,
		"name":   item.Name,
		"amount": item.Amount,
		"status": item.Status,
	}
}
