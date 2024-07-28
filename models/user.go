package models

import "github.com/nanda03dev/gnosql_client"

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
