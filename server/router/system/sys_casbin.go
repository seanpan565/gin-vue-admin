package system

// InitCasbinRouter 注册Casbin权限策略相关路由

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	{ // 需操作记录中间件的路由
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{ // 无需操作记录中间件的查询路由
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
