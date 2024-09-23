package common

type AppResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewAppResponse(data, paging, filter interface{}) *AppResponse {
	return &AppResponse{Data: data, Paging: paging, Filter: filter}
}

func SimplyAppResponse(data interface{}) *AppResponse {
	return &AppResponse{Data: data, Paging: nil, Filter: nil}
}
