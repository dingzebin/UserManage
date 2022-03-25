package util

type Pager struct {
	PageSize   int                    `json:"pageSize"`
	PageNum    int                    `json:"pageNum"`
	Parameters map[string]interface{} `json:"parameters"`
}

func (p *Pager) Start() int {
	return (p.PageNum - 1) * p.PageSize
}

func (p *Pager) Limit() int {
	return (p.PageNum-1)*p.PageSize + p.PageSize
}

type PageResult struct {
	Total   int
	Records interface{}
}
