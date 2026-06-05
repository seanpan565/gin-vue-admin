// Package announcement 公告管理插件，提供公告 CRUD 与前台展示能力。
package announcement

// plugin.go 插件入口，安装时初始化 API、菜单、字典、数据库与路由。

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

// Plugin 全局插件实例。
var Plugin = new(plugin)

// plugin 公告插件实现。
type plugin struct{}

// init 向插件框架注册本插件。
func init() {
	interfaces.Register(Plugin)
}

// Register 安装插件时依次初始化各子模块。
func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	// 如果需要配置文件，请到config.Config中填充配置结构，且到下方发放中填入其在config.yaml中的key
	// initialize.Viper()
	// 安装插件时候自动注册的api数据请到下方法.Api方法中实现
	initialize.Api(ctx)
	// 安装插件时候自动注册的Menu数据请到下方法.Menu方法中实现
	initialize.Menu(ctx)
	// 安装插件时候自动注册的Dictionary数据请到下方法.Dictionary方法中实现
	initialize.Dictionary(ctx)
	initialize.Gorm(ctx)
	initialize.Router(group)
}
