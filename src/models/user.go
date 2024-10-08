package models

import (
	"encoding/json"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gnosql_client"
)

type User struct {
	DocId   string `json:"docId" bson:"docId"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Address string `json:"address" bson:"address"`
	CityID  string `json:"cityId" bson:"cityId"`
}

var UserGnosql = gnosql_client.CollectionInput{
	CollectionName: "users",
	IndexKeys:      []string{"cityId"},
}

func (user User) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":   user.DocId,
		"name":    user.Name,
		"email":   user.Email,
		"address": user.Address,
		"cityId":  user.CityID,
	}
}

func (user User) ToModel(entityDocument gnosql_client.Document) User {
	entityString, _ := json.Marshal(entityDocument)

	var parsedEntity User
	json.Unmarshal(entityString, &parsedEntity)

	return parsedEntity

	// return User{
	// 	DocId:   GetStringValue(userDocument, "docId"),
	// 	Name:    GetStringValue(userDocument, "name"),
	// 	Email:   GetStringValue(userDocument, "email"),
	// 	Address: GetStringValue(userDocument, "address"),
	// 	CityID:  GetStringValue(userDocument, "cityId"),
	// }
}

func (user User) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      user.DocId,
		EntityType:    global_constant.ENTITY_USER,
		OperationType: operationType,
	}
}

func (user User) ToUpdatedDocument(newUser User) User {
	userDocument := user.ToDocument()
	newUserDocument := newUser.ToDocument()

	for key, value := range newUserDocument {
		if value != nil && value != "" {
			userDocument[key] = value
		}
	}

	return user.ToModel(userDocument)
}
