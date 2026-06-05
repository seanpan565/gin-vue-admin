// Package mcpTool MCP 分页响应通用结构。
package mcpTool

// pageResultData 通用分页数据结构。
type pageResultData[T any] struct {
	List     T     `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
