package main

var (
	DefaultPage     int64 = 1
	DefaultPageSize int64 = 5
)

type PageInfo struct {
	Page     int64 `json:"page" form:"page" binding:"required" default:"1"`
	PageSize int64 `json:"page_size"  form:"page_size"  binding:"omitempty" default:"5"`
}

type ResCommonStruct struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type ResPageStruct struct {
	Count  int64 `json:"count"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ListCommonStruct struct {
	Total       int64   `json:"total"`
	TotalPage   float64 `json:"total_page"`
	CurrentPage int64   `json:"current_page"`
}
