package models

import (
	"encoding/json"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gnosql_client"
)

type Item struct {
	DocId      string            `json:"docId" bson:"docId"`
	OrderId    string            `json:"orderId" bson:"orderId"`
	ProductId  string            `json:"productId" bson:"productId"`
	Amount     int               `json:"amount" bson:"amount"`
	StatusCode common.StatusCode `json:"statusCode" bson:"statusCode"`
}

var ItemGnosql = gnosql_client.CollectionInput{
	CollectionName: "items",
	IndexKeys:      []string{},
}

func (item Item) ToModel(entityDocument gnosql_client.Document) Item {
	entityString, _ := json.Marshal(entityDocument)

	var parsedEntity Item
	json.Unmarshal(entityString, &parsedEntity)

	return parsedEntity

	// return Item{
	// 	DocId:      GetStringValue(itemDocument, "docId"),
	// 	OrderId:    GetStringValue(itemDocument, "orderId"),
	// 	ProductId:  GetStringValue(itemDocument, "productId"),
	// 	Amount:     GetIntegerValue(itemDocument, "amount"),
	// 	StatusCode: GetValue[common.StatusCode](itemDocument, "statusCode"),
	// }
}

func (item Item) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":      item.DocId,
		"orderId":    item.OrderId,
		"productId":  item.ProductId,
		"amount":     item.Amount,
		"statusCode": item.StatusCode,
	}
}

func (item Item) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      item.DocId,
		EntityType:    global_constant.ENTITY_ITEM,
		OperationType: operationType,
	}
}

func (item Item) ToUpdatedDocument(newItem Item) Item {
	itemDocument := item.ToDocument()
	newItemDocument := newItem.ToDocument()

	// statusCode should not be updated
	newItemDocument["statusCode"] = itemDocument["statusCode"]

	for key, value := range newItemDocument {
		if value != nil && value != "" {
			itemDocument[key] = value
		}
	}

	return item.ToModel(itemDocument)
}
