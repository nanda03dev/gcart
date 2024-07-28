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

type SuccessCodeType struct {
	ORDER_INITIATED   string
	ORDER_TIMEOUT     string
	PAYMENT_INITIATED string
}

type ErrorCodeType struct {
	ORDER_TIMEOUT   string
	PAYMENT_TIMEOUT string
}

type EntityNameType string
type OperationNameType string

type EntitiesType struct {
	Order   EntityNameType
	Item    EntityNameType
	Payment EntityNameType
	Product EntityNameType
	Event   EntityNameType
}

type OperationsType struct {
	Create OperationNameType
	Update OperationNameType
	Delete OperationNameType
}

type EventType struct {
	EntityId      string
	EntityType    EntityNameType
	OperationType OperationNameType
	RetryCount    int
}

type GnoSQLCollectionSchemaType struct {
	CollectionName string
	IndexKeys      []string
}
