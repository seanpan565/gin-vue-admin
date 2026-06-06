package system

// InitSysErrorRouter 注册错误日志相关路由

import (
	"mall-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysErrorRouter struct{}

// InitSysErrorRouter 初始化 错误日志 路由信息
func (s *SysErrorRouter) InitSysErrorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
    sysErrorRouter := Router.Group("sysError").Use(middleware.OperationRecord())
    sysErrorRouterWithoutRecord := Router.Group("sysError")
    sysErrorRouterWithoutAuth := PublicRouter.Group("sysError")
    { // 需操作记录中间件的路由
        sysErrorRouter.DELETE("deleteSysError", sysErrorApi.DeleteSysError)           // 删除错误日志
        sysErrorRouter.DELETE("deleteSysErrorByIds", sysErrorApi.DeleteSysErrorByIds) // 批量删除错误日志
        sysErrorRouter.PUT("updateSysError", sysErrorApi.UpdateSysError)              // 更新错误日志
        sysErrorRouter.GET("getSysErrorSolution", sysErrorApi.GetSysErrorSolution)    // 触发错误日志处理
    }
    { // 无需操作记录中间件的查询路由
        sysErrorRouterWithoutRecord.GET("findSysError", sysErrorApi.FindSysError)       // 根据ID获取错误日志
        sysErrorRouterWithoutRecord.GET("getSysErrorList", sysErrorApi.GetSysErrorList) // 获取错误日志列表
    }
    { // 公开路由（前端上报错误：限流 + 可选密钥）
        sysErrorRouterWithoutAuth.POST("createSysError", middleware.AuthRateLimit(), middleware.ErrorReportGuard(), sysErrorApi.CreateSysError)
    }
}
