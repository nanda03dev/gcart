package models

import "github.com/nanda03dev/gnosql_client"

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
