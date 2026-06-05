// Example模块Service层：示例演示业务逻辑
package example

// ServiceGroup 示例模块Service聚合
type ServiceGroup struct {
	CustomerService

	AttachmentCategoryService
	FileUploadAndDownloadService
}
