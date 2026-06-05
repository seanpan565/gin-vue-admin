// Package plugin 定义 v1 插件接口，用于路由级插件扩展。
package plugin

// v1 插件模式接口，插件自行注册路由并返回路径前缀。
import (
	"github.com/gin-gonic/gin"
)

const (
	OnlyFuncName = "Plugin"
)

// Plugin 插件模式接口化
type Plugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)

	// RouterPath 用户返回注册路由
	RouterPath() string
}
