// Package mall 商城后台管理 API
package mall

import "mall-admin/server/service"

// ApiGroup 商城后台 API 聚合。
type ApiGroup struct {
	MemberApi
}

var memberService = service.ServiceGroupApp.MallServiceGroup.MemberService
