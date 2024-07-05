package global_constant

type StatusCodes struct {
	ORDER_INITIATED string
	ORDER_TIMEOUT   string
}

var ORDER_SUCCESS_STATUS_CODE = StatusCodes{
	ORDER_INITIATED: "3201",
}

var ORDER_ERROR_STATUS_CODE = StatusCodes{
	ORDER_TIMEOUT: "3408",
}
