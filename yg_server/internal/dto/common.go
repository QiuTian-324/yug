package dto

// 分页响应
type PaginationResponse struct {
	Total int         `json:"total"` // 总数
	Items interface{} `json:"items"` // 数据
}

// 分页请求
type PaginationRequest struct {
	Page     int `json:"page"`      // 页码
	PageSize int `json:"page_size"` // 每页数量
}

// 分页
func (p *PaginationRequest) Paginate() (int, int) {
	return (p.Page - 1) * p.PageSize, p.PageSize
}
