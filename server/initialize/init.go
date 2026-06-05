// Package initialize 应用启动初始化层，负责数据库、路由、插件、定时任务等组件装配。
// init.go 注册系统重载等全局事件处理函数。
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// 初始化全局函数
func SetupHandlers() {
	// 注册系统重载处理函数
	utils.GlobalSystemEvents.RegisterReloadHandler(func() error {
		return Reload()
	})
}
