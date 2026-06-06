// Package admin 管理后台 API 层：HTTP 请求处理、参数校验、调用 Service。
package admin

import (
	"mall-admin/server/api/admin/example"
	"mall-admin/server/api/admin/mall"
	"mall-admin/server/api/admin/system"
)

// ApiGroup 聚合系统与示例模块 API
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	MallApiGroup    mall.ApiGroup
}
