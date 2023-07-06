package response

type PageData struct {
	TotalCount int64       `json:"totalCount"`
	Data       interface{} `json:"data"`
}

func NewPageData(count int64, data interface{}) PageData {
	return PageData{TotalCount: count, Data: data}
}
