// Package initialize 公告插件安装初始化。
package initialize

// dictionary.go 安装时注册公告字典（暂无项）。

import (
	"context"
	model "mall-admin/server/model/system"
	"mall-admin/server/plugin/plugin-tool/utils"
)

// Dictionary 注册公告字典数据。
func Dictionary(ctx context.Context) {
	entities := []model.SysDictionary{}
	utils.RegisterDictionaries(entities...)
}
