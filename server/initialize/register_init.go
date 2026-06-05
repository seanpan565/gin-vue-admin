// register_init.go 侧载 source 包，触发各数据源 init 注册。
package initialize

import (
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/example"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
