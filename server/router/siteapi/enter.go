// C 端站点路由组
package siteapi

import (
	api "mall-admin/server/api/siteapi"
)

type RouterGroup struct {
	MemberAuthRouter
}

var (
	authApi = api.ApiGroupApp.AuthApiGroup.AuthApi
)
