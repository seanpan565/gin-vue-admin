package system

// InitAutoCodeRouter 注册代码生成相关路由

import "github.com/gin-gonic/gin"

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	publicAutoCodeRouter := RouterPublic.Group("autoCode")
	{ // 需鉴权的数据库元信息查询路由
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables)
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn)
	}
	{ // 代码模板预览与生成路由
		autoCodeRouter.POST("preview", autoCodeTemplateApi.Preview)
		autoCodeRouter.POST("createTemp", autoCodeTemplateApi.Create)
		autoCodeRouter.POST("addFunc", autoCodeTemplateApi.AddFunc)
	}
	{ // MCP服务管理路由
		autoCodeRouter.POST("mcp", autoCodeTemplateApi.MCP)
		autoCodeRouter.POST("mcpStatus", autoCodeTemplateApi.MCPStatus)
		autoCodeRouter.POST("mcpStart", autoCodeTemplateApi.MCPStart)
		autoCodeRouter.POST("mcpStop", autoCodeTemplateApi.MCPStop)
		autoCodeRouter.POST("mcpList", autoCodeTemplateApi.MCPList)
		autoCodeRouter.POST("mcpRoutes", autoCodeTemplateApi.MCPRoutes)
		autoCodeRouter.POST("mcpTest", autoCodeTemplateApi.MCPTest)
	}
	{ // 代码包与AI工作流会话路由
		autoCodeRouter.POST("getPackage", autoCodePackageApi.All)
		autoCodeRouter.POST("delPackage", autoCodePackageApi.Delete)
		autoCodeRouter.POST("createPackage", autoCodePackageApi.Create)
		autoCodeRouter.POST("saveAIWorkflowSession", aiWorkflowSessionApi.Save)
		autoCodeRouter.POST("getAIWorkflowSessionList", aiWorkflowSessionApi.GetList)
		autoCodeRouter.POST("getAIWorkflowSessionDetail", aiWorkflowSessionApi.GetDetail)
		autoCodeRouter.POST("deleteAIWorkflowSession", aiWorkflowSessionApi.Delete)
		autoCodeRouter.POST("dumpAIWorkflowMarkdown", aiWorkflowSessionApi.DumpMarkdown)
	}
	{ // 代码模板列表路由
		autoCodeRouter.GET("getTemplates", autoCodePackageApi.Templates)
	}
	{ // 插件打包与安装路由
		autoCodeRouter.POST("pubPlug", autoCodePluginApi.Packaged)
		autoCodeRouter.POST("installPlugin", autoCodePluginApi.Install)
		autoCodeRouter.POST("removePlugin", autoCodePluginApi.Remove)
		autoCodeRouter.GET("getPluginList", autoCodePluginApi.GetPluginList)
	}
	{ // 公开路由（LLM自动生成与插件初始化，无需鉴权）
		publicAutoCodeRouter.POST("llmAuto", autoCodeApi.LLMAuto)
		publicAutoCodeRouter.POST("llmAutoSSE", autoCodeApi.LLMAutoSSE)
		publicAutoCodeRouter.POST("initMenu", autoCodePluginApi.InitMenu)
		publicAutoCodeRouter.POST("initAPI", autoCodePluginApi.InitAPI)
		publicAutoCodeRouter.POST("initDictionary", autoCodePluginApi.InitDictionary)
	}
}
