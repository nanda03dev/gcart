package models

import (
	"github.com/nanda03dev/go2ms/common"
)

type User struct {
	DocId   string `json:"docId" bson:"docId"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Address string `json:"address" bson:"address"`
	CityID  string `json:"cityId" bson:"cityId"`
}

var UserGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "users",
	IndexKeys:      []string{"cityId"},
}
