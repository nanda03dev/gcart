package constant

type ErrorCode map[string]int

var ORDER_SUCCESS_STATUS_CODE = ErrorCode{
	"ORDER_TIMEOUT": 3201,
}
var ORDER_ERROR_STATUS_CODE = ErrorCode{
	"ORDER_TIMEOUT": 3408,
}
