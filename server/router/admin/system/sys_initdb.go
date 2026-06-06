package system

// InitInitRouter 注册数据库初始化相关路由

import (
	"mall-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init").Use(middleware.AuthRateLimit())
	{
		initRouter.POST("initdb", dbApi.InitDB)
		initRouter.POST("checkdb", dbApi.CheckDB)
	}
}
