package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/nanda03dev/gnosql_client"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
)

func Generate16DigitUUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}

// extractTimestampFromUUID extracts the timestamp from a version 1 UUID
func ExtractTimestampFromUUID(uuidStr string) time.Time {
	u, err := uuid.Parse(uuidStr)
	if err != nil {
		print(err)
	}
	// Version 1 UUID layout: time_low-time_mid-time_hi_and_version-clock_seq_hi_and_reserved-clock_seq_low-node
	// Extract timestamp from time_low, time_mid, and time_hi_and_version
	timestamp := int64(u[0])<<56 | int64(u[1])<<48 | int64(u[2])<<40 | int64(u[3])<<32 | int64(u[4])<<24 | int64(u[5])<<16 | int64(u[6])<<8 | int64(u[7])
	return time.Unix(0, timestamp)
}

func GetGnosqlCollection(entityType common.EntityNameType) gnosql_client.CollectionInput {
	switch entityType {
	case global_constant.ENTITY_CITY:
		return models.CityGnosql
	case global_constant.ENTITY_USER:
		return models.UserGnosql
	case global_constant.ENTITY_PRODUCT:
		return models.ProductGnosql
	case global_constant.ENTITY_ORDER:
		return models.OrderGnosql
	case global_constant.ENTITY_ITEM:
		return models.ItemGnosql
	case global_constant.ENTITY_PAYMENT:
		return models.PaymentGnosql

	default:
		return gnosql_client.CollectionInput{}
	}
}

func IsRequireToStoreEvent(entityType common.EntityNameType) bool {
	switch entityType {
	case global_constant.ENTITY_CITY,
		global_constant.ENTITY_USER,
		global_constant.ENTITY_PRODUCT,
		global_constant.ENTITY_ITEM:
		return false
	case global_constant.ENTITY_ORDER,
		global_constant.ENTITY_PAYMENT:
		return true
	default:
		return false
	}
}

func IsEventTimeExpired(entityType common.EntityNameType, eventCreatedAt time.Time) bool {
	var expireTime time.Time

	switch entityType {
	case global_constant.ENTITY_CITY,
		global_constant.ENTITY_USER,
		global_constant.ENTITY_PRODUCT,
		global_constant.ENTITY_ITEM:
		return false
	case global_constant.ENTITY_ORDER:
		expireTime = eventCreatedAt.Add(30 * time.Second)
	case global_constant.ENTITY_PAYMENT:
		expireTime = eventCreatedAt.Add(30 * time.Second)
	}

	return time.Now().After(expireTime)
}
