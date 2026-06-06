// Package example 文件上传与媒体库 API
package example

import "mall-admin/server/service"

// ApiGroup 示例模块 API 分组（仅保留上传相关）。
type ApiGroup struct {
	AttachmentCategoryApi
	FileUploadAndDownloadApi
}

var (
	attachmentCategoryService    = service.ServiceGroupApp.ExampleServiceGroup.AttachmentCategoryService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
