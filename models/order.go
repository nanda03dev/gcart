package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID      primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Amount  int                  `json:"amount" bson:"amount"`
	UserID  primitive.ObjectID   `json:"userId" bson:"userID"`
	Status  bool                 `bson:"status"`
	ItemsID []primitive.ObjectID `json:"itemIds" bson:"itemIds"`
}
