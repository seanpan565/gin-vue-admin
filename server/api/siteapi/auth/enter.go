// Package auth C 端会员认证 API
package auth

import "mall-admin/server/service"

// ApiGroup C 端认证 API 聚合。
type ApiGroup struct {
	AuthApi
}

var memberService = service.ServiceGroupApp.MallServiceGroup.MemberService
