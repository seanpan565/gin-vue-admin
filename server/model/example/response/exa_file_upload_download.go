package response

import "mall-admin/server/model/example"

// 文件上传下载相关响应结构
type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
