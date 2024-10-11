package restaurantratingmodel

type Filter struct {
	Point   int    `json:"point,omitempty" form:"column:point"`
	Comment string `json:"comment,omitempty" form:"column:comment"`
}

func (f *Filter) FullFill() {
	if f.Point <= 0 {
		f.Point = 1
	}
}
