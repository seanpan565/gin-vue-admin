package system

// InitSysVersionRouter 注册版本管理相关路由

import (
	"mall-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysVersionRouter struct{}

// InitSysVersionRouter 初始化 版本管理 路由信息
func (s *SysVersionRouter) InitSysVersionRouter(Router *gin.RouterGroup) {
	sysVersionRouter := Router.Group("sysVersion").Use(middleware.OperationRecord())
	sysVersionRouterWithoutRecord := Router.Group("sysVersion")
	{ // 需操作记录中间件的路由
		sysVersionRouter.DELETE("deleteSysVersion", sysVersionApi.DeleteSysVersion)           // 删除版本管理
		sysVersionRouter.DELETE("deleteSysVersionByIds", sysVersionApi.DeleteSysVersionByIds) // 批量删除版本管理
		sysVersionRouter.POST("exportVersion", sysVersionApi.ExportVersion)                   // 导出版本数据
		sysVersionRouter.POST("importVersion", sysVersionApi.ImportVersion)                   // 导入版本数据
	}
	{ // 无需操作记录中间件的查询路由
		sysVersionRouterWithoutRecord.GET("findSysVersion", sysVersionApi.FindSysVersion)           // 根据ID获取版本管理
		sysVersionRouterWithoutRecord.GET("getSysVersionList", sysVersionApi.GetSysVersionList)     // 获取版本管理列表
		sysVersionRouterWithoutRecord.GET("downloadVersionJson", sysVersionApi.DownloadVersionJson) // 下载版本JSON数据
	}
}
