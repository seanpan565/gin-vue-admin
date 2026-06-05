package system

// InitBaseRouter 注册登录、验证码等基础路由

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
