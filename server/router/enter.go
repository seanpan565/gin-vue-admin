// 路由层：注册HTTP路由、挂载中间件、绑定API处理函数
package router

import (
	"mall-admin/server/router/admin"
	"mall-admin/server/router/siteapi"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Admin   admin.RouterGroup
	SiteAPI siteapi.RouterGroup
}
