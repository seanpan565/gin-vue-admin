// Package response 定义示例模块 API 响应结构。
package response

import "mall-admin/server/model/example"

// 断点续传相关响应结构
type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}
