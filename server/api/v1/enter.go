// Package v1 API层：HTTP请求处理，参数绑定校验，调用Service，统一响应
package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

// ApiGroup 聚合系统与示例模块API
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}
