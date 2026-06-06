// Package siteapi C 端站点 API 层：面向商城消费者的 HTTP 接口。
package siteapi

import (
	"mall-admin/server/api/siteapi/auth"
)

// ApiGroup 聚合 C 端各业务模块 API。
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	AuthApiGroup auth.ApiGroup
}
