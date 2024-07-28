package models

import (
	"time"

	"github.com/nanda03dev/go2ms/common"
)

type Event struct {
	DocId         string                   `json:"docId" bson:"docId"`
	EntityId      string                   `json:"entityId" bson:"entityId"`
	EntityType    common.EntityNameType    `json:"entityType" bson:"entityType"`
	OperationType common.OperationNameType `json:"operationType" bson:"operationType"`
	ExpireTime    time.Time                `json:"expireTime" bson:"expireTime"`
}
