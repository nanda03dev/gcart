package models

import (
	"encoding/json"

	"github.com/nanda03dev/gcart/common"
	"github.com/nanda03dev/gcart/global_constant"
	"github.com/nanda03dev/gnosql_client"
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

func (city City) ToModel(cityDocument gnosql_client.Document) City {
	cityString, _ := json.Marshal(cityDocument)

	var parsedEntity City
	json.Unmarshal(cityString, &parsedEntity)

	return parsedEntity
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

func (city City) ToUpdatedDocument(newCity City) City {
	cityDocument := city.ToDocument()
	newCityDocument := newCity.ToDocument()

	for key, value := range newCityDocument {
		if value != nil && value != "" {
			cityDocument[key] = value
		}
	}

	return city.ToModel(cityDocument)
}
