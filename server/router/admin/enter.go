// 管理后台路由组聚合
package admin

import (
	"mall-admin/server/router/admin/example"
	"mall-admin/server/router/admin/mall"
	"mall-admin/server/router/admin/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Mall    mall.RouterGroup
}
