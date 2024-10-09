package restaurantratingmodel

type Filter struct {
	Point   int    `json:"point,omitempty" gorm:"column:point,omitempty"`
	Comment string `json:"commet,omitempty" gorm:"column:commet,omitempty"`
}

func (f *Filter) FullFill() {
	if f.Point <= 0 {
		f.Point = 1
	}
}
