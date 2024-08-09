package global_constant

import (
	"github.com/nanda03dev/gcart/common"
)

const (
	ENTITY_CITY           common.EntityNameType = "CITY"
	ENTITY_USER           common.EntityNameType = "USER"
	ENTITY_ORDER          common.EntityNameType = "ORDER"
	ENTITY_ITEM           common.EntityNameType = "ITEM"
	ENTITY_PAYMENT        common.EntityNameType = "PAYMENT"
	ENTITY_REFUND_PAYMENT common.EntityNameType = "REFUND_PAYMENT"
	ENTITY_PRODUCT        common.EntityNameType = "PRODUCT"
	ENTITY_EVENT          common.EntityNameType = "EVENT"
)

const (
	OPERATION_CREATE    common.OperationType = "CREATE"
	OPERATION_UPDATE    common.OperationType = "UPDATE"
	OPERATION_DELETE    common.OperationType = "DELETE"
	OPERATION_CONFIRMED common.OperationType = "CONFIRMED"
	OPERATION_CANCELLED common.OperationType = "CANCELLED"
)
const (
	CHECK_TIMEOUT      common.CheckProcess = "CHECK_TIMEOUT"
	CHECK_TIMEOUT_DONE common.CheckProcess = "CHECK_TIMEOUT_DONE"
)

const (
	// Order status statusCode
	ORDER_INITIATED common.StatusCode = "3201"
	ORDER_CONFIRMED common.StatusCode = "3202"
	ORDER_CANCELLED common.StatusCode = "3203"
	ORDER_TIMEOUT   common.StatusCode = "3408"

	// Order status statusCode
	ITEM_INITIATED common.StatusCode = "4201"
	ITEM_CONFIRMED common.StatusCode = "4202"
	ITEM_CANCELLED common.StatusCode = "4203"
	ITEM_TIMEOUT   common.StatusCode = "4408"

	// Payment status statusCode
	PAYMENT_INITIATED common.StatusCode = "5201"
	PAYMENT_CONFIRMED common.StatusCode = "5202"
	PAYMENT_CANCELLED common.StatusCode = "5203"
	PAYMENT_TIMEOUT   common.StatusCode = "5408"

	// Refund Payment status statusCode
	REFUND_PAYMENT_INITIATED common.StatusCode = "6201"
	REFUND_PAYMENT_CONFIRMED common.StatusCode = "6202"
	REFUND_PAYMENT_CANCELLED common.StatusCode = "6203"
	REFUND_PAYMENT_TIMEOUT   common.StatusCode = "6408"
)

const (
	ENTITY_CREATED_SUCCESSFULLY   = "ENTITY_CREATED_SUCCESSFULLY"
	ENTITY_FETCHED_SUCCESSFULLY   = "ENTITY_FETCHED_SUCCESSFULLY"
	ENTITY_UPDATED_SUCCESSFULLY   = "ENTITY_UPDATED_SUCCESSFULLY"
	ENTITY_DELETED_SUCCESSFULLY   = "ENTITY_DELETED_SUCCESSFULLY"
	ENTITY_CONFIRMED_SUCCESSFULLY = "ENTITY_CONFIRMED_SUCCESSFULLY"

	ERROR_WHILE_PROCESSING                                    = "ERROR_WHILE_PROCESSING"
	ERROR_ENTITY_NOT_FOUND                                    = "ERROR_ENTITY_NOT_FOUND"
	ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_PAYMENT_PENDING    = "ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_PAYMENT_PENDING"
	ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_ITEM_CONFIRM_ISSUE = "ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_ITEM_CONFIRM_ISSUE"
	ERROR_ENTITY_ALREADY_COMPLETED_OR_TIMEDOUT                = "ERROR_ENTITY_ALREADY_COMPLETED_OR_TIMEDOUT"
)
