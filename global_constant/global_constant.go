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
	CHECK_TIMEOUT      common.CheckProcess = "CHECK_TIMEOUT"
	CHECK_TIMEOUT_DONE common.CheckProcess = "CHECK_TIMEOUT_DONE"
)

const (
	// Order status statusCode
	ORDER_INITIATED common.StatusCode = "3201"
	ORDER_TIMEOUT   common.StatusCode = "3408"

	// Order status statusCode
	ITEM_INITIATED common.StatusCode = "4201"
	ITEM_TIMEOUT   common.StatusCode = "4408"

	// Payment status statusCode
	PAYMENT_INITIATED common.StatusCode = "5201"
	PAYMENT_TIMEOUT   common.StatusCode = "5408"
)

const (
	DATA_CREATED_SUCCESSFULLY            = "DATA_CREATED_SUCCESSFULLY"
	DATA_FETCHED_SUCCESSFULLY            = "DATA_FETCHED_SUCCESSFULLY"
	DATA_UPDATED_SUCCESSFULLY            = "DATA_UPDATED_SUCCESSFULLY"
	DATA_DELETED_SUCCESSFULLY            = "DATA_DELETED_SUCCESSFULLY"
	ERROR_WHILE_PROCESSING               = "ERROR_WHILE_PROCESSING"
	ENTITY_NOT_FOUND                     = "ENTITY_NOT_FOUND"
	ENTITY_ALREADY_COMPLETED_OR_TIMEDOUT = "ENTITY_ALREADY_COMPLETED_OR_TIMEDOUT"
)
