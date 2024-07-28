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
	case global_constant.Entities.City:
		return models.CityGnosql
	case global_constant.Entities.Product:
		return models.ProductGnosql
	case global_constant.Entities.Order:
		return models.OrderGnosql
	case global_constant.Entities.Item:
		return models.ItemGnosql
	case global_constant.Entities.Payment:
		return models.PaymentGnosql
	default:
		return gnosql_client.CollectionInput{}
	}
}
