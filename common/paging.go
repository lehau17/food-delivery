package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
}

func (p *Paging) FullFill() {
	if p.Limit <= 0 {
		p.Limit = 50
	}
	if p.Page <= 0 {
		p.Page = 1
	}
}

type PagingCursor struct {
	Limit int    `json:"limit" form:"limit"`
	Total int64  `json:"total" form:"total"`
	Curor string `json:"curor" form:"curor"`
}
