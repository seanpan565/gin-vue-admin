// Package example 示例业务模块API
package example

import "mall-admin/server/service"

// ApiGroup 示例模块API分组
type ApiGroup struct {
	CustomerApi

	AttachmentCategoryApi
	FileUploadAndDownloadApi
}

// 注入 service 层依赖
var (
	customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService

	attachmentCategoryService    = service.ServiceGroupApp.ExampleServiceGroup.AttachmentCategoryService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
