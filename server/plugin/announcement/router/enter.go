// Package router 公告插件路由层。
package router

// enter.go 聚合公告路由及 API 引用。
import "mall-admin/server/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

// router 公告路由分组。
type router struct{ Info info }
