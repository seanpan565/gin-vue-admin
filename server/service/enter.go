// Service层：业务逻辑，不依赖gin.Context，通过GORM操作数据库
package service

// ServiceGroup 聚合各子模块Service
import (
	"mall-admin/server/service/example"
	"mall-admin/server/service/mall"
	"mall-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	MallServiceGroup    mall.ServiceGroup
}
