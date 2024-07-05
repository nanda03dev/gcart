package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Amount int                `json:"amount" bson:"amount"`
	Status bool               `bson:"status"`
}
