package common

type Filter struct {
	Key   string
	Value interface{}
}

type Filters []Filter

type RequestFilterBody struct {
	ListOfFilter Filters `json:"filters"`
	Size         int     `json:"size"`
	SortBody     Sort    `json:"sort"`
}

type Sort struct {
	Key   string `json:"Key"`
	Order int    `json:"order"`
}
