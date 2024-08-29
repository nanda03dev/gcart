package models

import (
	"encoding/json"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gnosql_client"
)

type Product struct {
	DocId  string            `json:"docId" bson:"docId"`
	Name   string            `json:"name" bson:"name"`
	Amount int               `json:"amount" bson:"amount"`
	Status common.StatusCode `bson:"status"`
}

var ProductGnosql = gnosql_client.CollectionInput{
	CollectionName: "products",
	IndexKeys:      []string{},
}

func (product Product) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      product.DocId,
		"name":       product.Name,
		"amount":     product.Amount,
		"statusCode": product.Status,
	}
}

func (product Product) ToModel(entityDocument gnosql_client.Document) Product {
	entityString, _ := json.Marshal(entityDocument)

	var parsedEntity Product
	json.Unmarshal(entityString, &parsedEntity)

	return parsedEntity

	// return Product{
	// 	DocId:  GetStringValue(productDocument, "docId"),
	// 	Name:   GetStringValue(productDocument, "name"),
	// 	Amount: GetIntegerValue(productDocument, "amount"),
	// 	Status: GetValue[common.StatusCode](productDocument, "statusCode"),
	// }
}

func (product Product) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      product.DocId,
		EntityType:    global_constant.ENTITY_PRODUCT,
		OperationType: operationType,
	}
}

func (product Product) ToUpdatedDocument(newProduct Product) Product {
	productDocument := product.ToDocument()
	newProductDocument := newProduct.ToDocument()

	for key, value := range newProductDocument {
		if value != nil && value != "" {
			productDocument[key] = value
		}
	}

	return product.ToModel(productDocument)
}
