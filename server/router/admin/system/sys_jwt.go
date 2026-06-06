package system

// InitJwtRouter 注册JWT黑名单相关路由

import (
	"github.com/gin-gonic/gin"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
