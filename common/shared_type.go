package common

type FilterBody struct {
	Key   string
	Value interface{}
}

type FiltersBody []FilterBody

type RequestFilterBody struct {
	ListOfFilter FiltersBody `json:"filters"`
	Size         int         `json:"size"`
	SortBody     SortBody    `json:"sort"`
}

type SortBody struct {
	Key   string `json:"Key"`
	Order int    `json:"order"`
}
