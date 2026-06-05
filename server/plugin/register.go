// Package plugin 插件注册入口，通过空白导入加载各内置子插件。
package plugin

// register.go 空白导入子插件，触发其 init 完成注册。

import (
	_ "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
	_ "github.com/flipped-aurora/gin-vue-admin/server/plugin/auto"
)
