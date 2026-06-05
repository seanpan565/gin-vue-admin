// Package api 公告插件 API 层。
package api

// enter.go 聚合公告 API 及 Service 引用。
import "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

// api 公告 API 分组。
type api struct{ Info info }
