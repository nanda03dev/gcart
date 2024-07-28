package global_constant

import (
	"github.com/nanda03dev/go2ms/common"
)

var Entities = common.EntitiesType{
	City:    "City",
	User:    "User",
	Order:   "Order",
	Item:    "Item",
	Payment: "Payment",
	Product: "Product",
	Event:   "Event",
}

var Operations = common.OperationsType{
	Create: "Create",
	Update: "Update",
	Delete: "Delete",
}

var OrderSuccessCode = common.SuccessCodeType{
	ORDER_INITIATED: "3201",
}

var OrderErrorCode = common.ErrorCodeType{
	ORDER_TIMEOUT: "3408",
}

var PaymentSuccessCode = common.SuccessCodeType{
	PAYMENT_INITIATED: "5201",
}

var PaymentErrorCode = common.ErrorCodeType{
	PAYMENT_TIMEOUT: "5408",
}
