package models

import (
	"time"

	"github.com/nanda03dev/go2ms/common"
)

type Event struct {
	DocId         string                `json:"docId" bson:"docId"`
	EntityId      string                `json:"entityId" bson:"entityId"`
	EntityType    common.EntityNameType `json:"entityType" bson:"entityType"`
	OperationType common.OperationType  `json:"operationType" bson:"operationType"`
	CheckProcess  common.CheckProcess   `json:"checkProcess" bson:"checkProcess"`
	Published     bool                  `json:"published" bson:"published"`
	CreatedAt     time.Time             `json:"createdAt" bson:"createdAt"`
}
