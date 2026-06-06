// Package initialize 公告插件安装初始化。
package initialize

// viper.go 从配置文件加载公告插件配置。

import (
	"fmt"
	"mall-admin/server/global"
	"mall-admin/server/plugin/announcement/plugin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Viper 解析 announcement 配置节。
func Viper() {
	err := global.GVA_VP.UnmarshalKey("announcement", &plugin.Config)
	if err != nil {
		err = errors.Wrap(err, "初始化配置文件失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
