// plugin_biz_v1.go v1 插件注册（RouterGroup 模式），如邮件插件。
package initialize

import (
	"mall-admin/server/global"
	"mall-admin/server/plugin/email"
	"mall-admin/server/utils/plugin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		path := Plugin[i].RouterPath()
		global.GVA_LOG.Info("插件注册开始", zap.String("path", path))
		PluginGroup := group.Group(path)
		Plugin[i].Register(PluginGroup)
		global.GVA_LOG.Info("插件注册成功", zap.String("path", path))
	}
}

func bizPluginV1(group ...*gin.RouterGroup) {
	private := group[0]
	public := group[1]
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(private, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
		global.GVA_CONFIG.Email.IsLoginAuth,
	))
	_ = public
}
