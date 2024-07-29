package models

import (
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
)

type City struct {
	DocId       string `json:"docId" bson:"docId"`
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"countryCode" bson:"countryCode"`
}

var CityGnosql = gnosql_client.CollectionInput{
	CollectionName: "cities",
	IndexKeys:      []string{"countryCode"},
}

func (city City) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":       city.DocId,
		"name":        city.Name,
		"countryCode": city.CountryCode,
	}
}

func (city City) ToEvent(operationType common.OperationType) common.EventType {
	return common.EventType{
		EntityId:      city.DocId,
		EntityType:    global_constant.ENTITY_CITY,
		OperationType: operationType,
	}
}
