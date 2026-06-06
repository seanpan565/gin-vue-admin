// router.go 注册 Gin 总路由，挂载中间件、Swagger 及系统/业务/插件路由组。
package initialize

import (
	"mall-admin/server/docs"
	"mall-admin/server/global"
	"mall-admin/server/middleware"
	"mall-admin/server/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routers 初始化 Gin 引擎，区分公开/私有路由组并注册全部业务路由。
func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.GinRecovery(true))
	Router.Use(middleware.SecurityHeaders())
	Router.Use(middleware.DefaultLimit())
	if global.GVA_CONFIG.Security.MetricsEnable {
		Router.Use(middleware.PrometheusMetrics())
	}
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	adminRouter := router.RouterGroupApp.Admin
	systemRouter := adminRouter.System
	mallRouter := adminRouter.Mall
	fileUploadRouter := adminRouter.Example

	RegisterSecureStatic(Router)

	if global.GVA_CONFIG.Security.SwaggerEnable {
		docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
		Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.GVA_LOG.Info("register swagger handler")
	}

	prefix := global.GVA_CONFIG.System.RouterPrefix
	PublicGroup := Router.Group(prefix)
	PrivateGroup := Router.Group(prefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	if global.GVA_CONFIG.Security.MetricsEnable {
		metricsGroup := Router.Group(prefix)
		metricsGroup.GET("/metrics", middleware.MetricsAuth(), middleware.PrometheusHandler())
	}

	{
		PublicGroup.GET("/health", HealthHandler)
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}

	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitSysVersionRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysParamsRouter(PrivateGroup, PublicGroup)
		systemRouter.InitSysErrorRouter(PrivateGroup, PublicGroup)
		systemRouter.InitLoginLogRouter(PrivateGroup)
		systemRouter.InitApiTokenRouter(PrivateGroup)
		fileUploadRouter.InitFileUploadAndDownloadRouter(PrivateGroup)
		fileUploadRouter.InitAttachmentCategoryRouterRouter(PrivateGroup)
		mallRouter.InitMemberRouter(PrivateGroup)
	}

	InstallPlugin(PrivateGroup, PublicGroup, Router)

	initBizRouter(PrivateGroup, PublicGroup)

	global.GVA_ROUTERS = Router.Routes()

	global.GVA_LOG.Info("router register success")
	return Router
}
