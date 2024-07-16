package models

import (
	"time"

	"github.com/nanda03dev/go2ms/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID            primitive.ObjectID       `json:"id" bson:"_id,omitempty"`
	EntityId      primitive.ObjectID       `json:"entityId" bson:"entityId"`
	EntityType    common.EntityNameType    `json:"entityType" bson:"entityType"`
	OperationType common.OperationNameType `json:"operationType" bson:"operationType"`
	ExpireTime    time.Time                `json:"expireTime" bson:"expireTime"`
}
