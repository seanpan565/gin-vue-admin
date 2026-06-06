// Package request 定义示例模块 API 请求参数结构。
package request

import (
	"mall-admin/server/model/common/request"
)

// 文件上传下载相关请求参数
type ExaAttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}
