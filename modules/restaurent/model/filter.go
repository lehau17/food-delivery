package restaurentmodel

type Filter struct {
	Status []int `json:"status" form:"status"`
}
