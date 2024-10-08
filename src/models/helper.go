package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gnosql_client"
)

func Generate16DigitUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func GetAllGnosqlCollections() []gnosql_client.CollectionInput {
	return []gnosql_client.CollectionInput{
		CityGnosql,
		UserGnosql,
		ProductGnosql,
		OrderGnosql,
		ItemGnosql,
		PaymentGnosql,
		RefundPaymentGnosql,
	}
}

func GetGnosqlCollection(entityType common.EntityNameType) gnosql_client.CollectionInput {
	switch entityType {
	case global_constant.ENTITY_CITY:
		return CityGnosql
	case global_constant.ENTITY_USER:
		return UserGnosql
	case global_constant.ENTITY_PRODUCT:
		return ProductGnosql
	case global_constant.ENTITY_ORDER:
		return OrderGnosql
	case global_constant.ENTITY_ITEM:
		return ItemGnosql
	case global_constant.ENTITY_PAYMENT:
		return PaymentGnosql
	case global_constant.ENTITY_REFUND_PAYMENT:
		return RefundPaymentGnosql

	default:
		return gnosql_client.CollectionInput{}
	}
}

func IsRequireToStoreEvent(entityType common.EntityNameType) bool {
	switch entityType {
	case
		global_constant.ENTITY_USER,
		global_constant.ENTITY_PRODUCT,
		global_constant.ENTITY_ITEM:
		return false
	case global_constant.ENTITY_ORDER,
		global_constant.ENTITY_PAYMENT,
		global_constant.ENTITY_REFUND_PAYMENT,
		global_constant.ENTITY_CITY:
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
		global_constant.ENTITY_ITEM,
		global_constant.ENTITY_REFUND_PAYMENT:
		return false
	case global_constant.ENTITY_ORDER:
		expireTime = eventCreatedAt.Add(2 * time.Minute)
	case global_constant.ENTITY_PAYMENT:
		expireTime = eventCreatedAt.Add(1 * time.Minute)
	}

	return time.Now().After(expireTime)
}

func GetCheckProcess(entityType common.EntityNameType, operationType common.OperationType) common.CheckProcess {
	switch entityType {
	case global_constant.ENTITY_ORDER, global_constant.ENTITY_PAYMENT:
		if operationType == global_constant.OPERATION_CREATE {
			return global_constant.CHECK_TIMEOUT
		}
	case global_constant.ENTITY_CITY,
		global_constant.ENTITY_USER,
		global_constant.ENTITY_PRODUCT,
		global_constant.ENTITY_REFUND_PAYMENT,
		global_constant.ENTITY_ITEM:
		return ""
	}
	return ""
}

func GetStringValue(document gnosql_client.Document, key string) string {
	value := document[key]
	return value.(string)
}
func GetIntegerValue(document gnosql_client.Document, key string) int {
	value := document[key]
	return value.(int)
}
func GetBoolValue(document gnosql_client.Document, key string) bool {
	value := document[key]
	return value.(bool)
}
func GetValue[T any](document gnosql_client.Document, key string) T {
	value := document[key]
	return value.(T)
}
