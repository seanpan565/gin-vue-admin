// Package auto 编程辅助插件，集成代码生成、MCP、技能管理等能力。
package auto

// plugin.go 插件入口，安装时初始化 API、菜单、字典、数据库与路由。

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/auto/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

// Plugin 全局插件实例，供框架 v2 插件系统加载。
var Plugin = new(plugin)

// plugin 编程辅助插件实现。
type plugin struct{}

// init 向插件框架注册本插件。
func init() {
	interfaces.Register(Plugin)
}

// Register 安装插件时依次初始化 API、菜单、字典、表结构与路由。
func (p *plugin) Register(engine *gin.Engine) {
	ctx := context.Background()
	initialize.Api(ctx)
	initialize.Menu(ctx)
	initialize.Dictionary(ctx)
	initialize.Gorm(ctx)
	initialize.Router(engine)
}
