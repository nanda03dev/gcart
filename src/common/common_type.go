package common

type FilterBodyType struct {
	Key   string
	Value interface{}
}

type FiltersBodyType []FilterBodyType

type SortBodyType struct {
	Key   string `json:"Key"`
	Order int    `json:"order"`
}

type RequestFilterBodyType struct {
	ListOfFilter FiltersBodyType `json:"filters"`
	Size         int             `json:"size"`
	SortBody     SortBodyType    `json:"sort"`
}

type OrderConfirmBody struct {
	OrderId string `json:"orderId"`
}

type PaymentConfirmBody struct {
	PaymentId string `json:"paymentId"`
}

type RefundPaymentConfirmBody struct {
	RefundPaymentId string `json:"refundPaymentId"`
}

type EntityNameType string
type OperationType string
type StatusCode string
type CheckProcess string

type EventType struct {
	EntityId      string
	EntityType    EntityNameType
	OperationType OperationType
	CheckProcess  CheckProcess
	RetryCount    int
}

type GnoSQLCollectionSchemaType struct {
	CollectionName string
	IndexKeys      []string
}

type SuccessResponse struct {
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

type ErrorResponse struct {
	Msg   any `json:"msg"`
	Error any `json:"error"`
}
