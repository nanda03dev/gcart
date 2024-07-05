package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderId    string             `json:"orderId" bson:"orderId"`
	ExpireTime time.Time          `json:"expireTime" bson:"expireTime"`
}
