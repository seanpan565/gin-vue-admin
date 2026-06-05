// 示例模块路由组：注册客户、附件分类、文件上传等示例业务路由
package example

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	CustomerRouter

	AttachmentCategoryRouter
	FileUploadAndDownloadRouter
}

var (
	exaCustomerApi = api.ApiGroupApp.ExampleApiGroup.CustomerApi

	attachmentCategoryApi       = api.ApiGroupApp.ExampleApiGroup.AttachmentCategoryApi
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
)
