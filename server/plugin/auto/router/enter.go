// Package router auto 插件路由层，挂载代码生成与技能管理接口。
package router

// enter.go 聚合 auto 插件路由分组及 API 引用。
import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

// RouterGroup auto 插件路由分组。
type RouterGroup struct {
	AutoCodeRouter
	SkillsRouter
}

var (
	autoCodeApi          = api.ApiGroupApp.SystemApiGroup.AutoCodeApi
	autoCodePluginApi    = api.ApiGroupApp.SystemApiGroup.AutoCodePluginApi
	autocodeHistoryApi   = api.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	autoCodePackageApi   = api.ApiGroupApp.SystemApiGroup.AutoCodePackageApi
	autoCodeTemplateApi  = api.ApiGroupApp.SystemApiGroup.AutoCodeTemplateApi
	skillsApi            = api.ApiGroupApp.SystemApiGroup.SkillsApi
	aiWorkflowSessionApi = api.ApiGroupApp.SystemApiGroup.AIWorkflowSessionApi
)

var RouterGroupApp = new(RouterGroup)
