// Package api 邮件插件 API 层。
package api

// ApiGroup 邮件 API 分组。
type ApiGroup struct {
	EmailApi
}

var ApiGroupApp = new(ApiGroup)
