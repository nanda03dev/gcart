package models

import (
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

func (city City) ToDocument() gnosql_client.Document {
	return gnosql_client.Document{
		"docId":       city.DocId,
		"name":        city.Name,
		"countryCode": city.CountryCode,
	}
}
