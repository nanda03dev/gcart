package models

import (
	"github.com/nanda03dev/go2ms/common"
)

type City struct {
	DocId       string `json:"docId" bson:"docId"`
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"countryCode" bson:"countryCode"`
}

var CityGnosql = common.GnoSQLCollectionSchemaType{
	CollectionName: "cities",
	IndexKeys:      []string{"countryCode"},
}
