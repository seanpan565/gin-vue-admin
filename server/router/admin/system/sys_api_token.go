package system

// InitApiTokenRouter 注册API Token相关路由

import (
	"mall-admin/server/api/admin"
	"mall-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiTokenRouter struct{}

func (s *ApiTokenRouter) InitApiTokenRouter(Router *gin.RouterGroup) {
	apiTokenRouter := Router.Group("sysApiToken").Use(middleware.OperationRecord())
	apiTokenApi := admin.ApiGroupApp.SystemApiGroup.ApiTokenApi
	{
		apiTokenRouter.POST("createApiToken", apiTokenApi.CreateApiToken)   // 签发Token
		apiTokenRouter.POST("getApiTokenList", apiTokenApi.GetApiTokenList) // 获取列表
		apiTokenRouter.POST("deleteApiToken", apiTokenApi.DeleteApiToken)   // 作废Token
	}
}
