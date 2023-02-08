package util

type Reponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
}

type ListResponse struct {
	Data  interface{} `json:"data"`
	Pagin Pagination  `json:"pagination"`
}

type Pagination struct {
	PageSize int   `json:"pageSize"`
	PageIdx  int   `json:"pageIdx"`
	Total    int64 `json:"total"`
}
