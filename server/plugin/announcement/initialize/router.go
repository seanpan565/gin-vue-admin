// Package initialize 公告插件安装初始化。
package initialize

// router.go 安装时挂载公告 HTTP 路由。
import (
	"mall-admin/server/global"
	"mall-admin/server/middleware"
	"mall-admin/server/plugin/announcement/router"
	"github.com/gin-gonic/gin"
)

// Router 注册公告路由入口。
func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.Info.Init(public, private)
}
