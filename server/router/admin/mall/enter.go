// 商城后台管理路由组
package mall

import (
	api "mall-admin/server/api/admin"
)

type RouterGroup struct {
	MemberRouter
}

var (
	memberApi = api.ApiGroupApp.MallApiGroup.MemberApi
)
