package global_constant

import (
	"github.com/nanda03dev/go2ms/common"
)

const (
	ENTITY_CITY    common.EntityNameType = "CITY"
	ENTITY_USER    common.EntityNameType = "USER"
	ENTITY_ORDER   common.EntityNameType = "ORDER"
	ENTITY_ITEM    common.EntityNameType = "ITEM"
	ENTITY_PAYMENT common.EntityNameType = "PAYMENT"
	ENTITY_PRODUCT common.EntityNameType = "PRODUCT"
	ENTITY_EVENT   common.EntityNameType = "EVENT"
)

const (
	OPERATION_CREATE common.OperationType = "CREATE"
	OPERATION_UPDATE common.OperationType = "UPDATE"
	OPERATION_DELETE common.OperationType = "DELETE"
)

const (
	// Order status code
	ORDER_INITIATED common.StatusCode = "3201"
	ORDER_TIMEOUT   common.StatusCode = "3408"

	// Payment status code
	PAYMENT_INITIATED common.StatusCode = "5201"
	PAYMENT_TIMEOUT   common.StatusCode = "5408"
)
