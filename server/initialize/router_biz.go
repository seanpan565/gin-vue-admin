// router_biz.go 业务自定义路由注册入口，在此扩展项目专属路由。
package initialize

import (
	"mall-admin/server/router"
	"github.com/gin-gonic/gin"
)

// initBizRouter 注册业务路由，开发者在此追加自定义 RouterGroup。
func initBizRouter(routers ...*gin.RouterGroup) {
	publicGroup := routers[1]

	siteAPIRouter := router.RouterGroupApp.SiteAPI

	// 商城 C 端 API：/site/auth/*（会员 token 走 x-member-token）
	siteAPIRouter.InitMemberAuthRouter(publicGroup.Group("site"))
}
