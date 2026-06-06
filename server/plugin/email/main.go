// Package email 邮件插件，提供 SMTP 发信与测试接口。
package email

// main.go 邮件插件入口，支持 v1 插件方式注册路由。

import (
	"mall-admin/server/plugin/email/global"
	"mall-admin/server/plugin/email/router"
	"github.com/gin-gonic/gin"
)

// emailPlugin 邮件插件实现。
type emailPlugin struct{}

// CreateEmailPlug 创建邮件插件并写入 SMTP 配置。
func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool, IsLoginAuth bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Nickname = Nickname
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	global.GlobalConfig.IsLoginAuth = IsLoginAuth
	return &emailPlugin{}
}

// Register 注册邮件相关 HTTP 路由。
func (*emailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
