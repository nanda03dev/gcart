package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount string             `json:"amount" bson:"amount"`
}
