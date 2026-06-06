package system

// InitSystemRouter 注册系统配置相关路由

import (
	"mall-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	sysRouterWithoutRecord := Router.Group("system")

	{ // 需操作记录中间件的路由
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
	}
	{ // 无需操作记录中间件的查询路由
		sysRouterWithoutRecord.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouterWithoutRecord.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
	}
}
