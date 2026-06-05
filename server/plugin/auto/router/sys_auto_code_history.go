// sys_auto_code_history.go 代码生成历史与回滚相关路由。
package router

import "github.com/gin-gonic/gin"

// AutoCodeHistoryRouter 代码生成历史路由组。
type AutoCodeHistoryRouter struct{}

// InitAutoCodeHistoryRouter 注册代码生成历史路由。
func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	{
		autoCodeHistoryRouter.POST("getMeta", autocodeHistoryApi.First)         // 根据 ID 获取元数据
		autoCodeHistoryRouter.POST("rollback", autocodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", autocodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", autocodeHistoryApi.GetList) // 获取回滚记录分页
	}
}
