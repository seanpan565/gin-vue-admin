// Package plugin 定义 v2 插件接口及注册机制。
package plugin

// v2 插件模式接口，插件直接向 Engine 注册路由。
import (
	"github.com/gin-gonic/gin"
)

// Plugin 插件模式接口化v2
type Plugin interface {
	// Register 注册路由
	Register(group *gin.Engine)
}
