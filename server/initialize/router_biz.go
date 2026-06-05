// router_biz.go 业务自定义路由注册入口，在此扩展项目专属路由。
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

// initBizRouter 注册业务路由，开发者在此追加自定义 RouterGroup。
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	holder(publicGroup, privateGroup)
}
