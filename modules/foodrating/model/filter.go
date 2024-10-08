package foodratingmodel

type Filter struct {
	Status  *int    `json:"status,omitempty" form:"status,omitempty"`
	Comment *string `json:"comment,omitempty" form:"comment,omitempty"`
}

func (f *Filter) FullFill() {
	if f.Status == nil {
		f.Status = new(int)
		*f.Status = 1
	}

}
