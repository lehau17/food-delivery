package foodmodel

import "log"

type Filter struct {
	Status *int `json:"status" form:"status,omitempty"`
}

func (f *Filter) FullFilter() {
	log.Println("Check status: ", f.Status)
	if f.Status == nil {
		f.Status = new(int)
		*f.Status = 1
	}
	log.Println("Check status: ", *f.Status)

}
