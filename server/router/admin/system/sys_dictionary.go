package system

// InitSysDictionaryRouter 注册字典管理相关路由

import (
	"mall-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	{ // 需操作记录中间件的路由
		sysDictionaryRouter.POST("createSysDictionary", dictionaryApi.CreateSysDictionary)   // 新建SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", dictionaryApi.DeleteSysDictionary) // 删除SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", dictionaryApi.UpdateSysDictionary)    // 更新SysDictionary
		sysDictionaryRouter.POST("importSysDictionary", dictionaryApi.ImportSysDictionary)   // 导入SysDictionary
		sysDictionaryRouter.GET("exportSysDictionary", dictionaryApi.ExportSysDictionary)    // 导出SysDictionary
	}
	{ // 无需操作记录中间件的查询路由
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", dictionaryApi.FindSysDictionary)       // 根据ID获取SysDictionary
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", dictionaryApi.GetSysDictionaryList) // 获取SysDictionary列表
	}
}
