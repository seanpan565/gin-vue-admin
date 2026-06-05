// Package initialize auto 插件安装初始化。
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/router"
	"github.com/gin-gonic/gin"
)

// router.go 安装时挂载 auto 插件 HTTP 路由。

// Router 注册 auto 插件路由入口。
func Router(engine *gin.Engine) {
	InitializeRouter(engine)
}

func InitializeRouter(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	router.RouterGroupApp.InitAutoCodeRouter(private, public)
	router.RouterGroupApp.InitAutoCodeHistoryRouter(private)
	router.RouterGroupApp.InitSkillsRouter(private, public)
}
