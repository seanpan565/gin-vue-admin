// 示例模块路由组：文件上传与媒体库分类
package example

import (
	api "mall-admin/server/api/admin"
)

type RouterGroup struct {
	AttachmentCategoryRouter
	FileUploadAndDownloadRouter
}

var (
	attachmentCategoryApi         = api.ApiGroupApp.ExampleApiGroup.AttachmentCategoryApi
	exaFileUploadAndDownloadApi   = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
)
