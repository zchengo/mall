package models

// Page 分页参数模型
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}
