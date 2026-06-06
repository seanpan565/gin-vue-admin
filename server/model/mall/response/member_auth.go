// member_auth.go C 端会员认证响应结构。
package response

import (
	"mall-admin/server/model/mall"
)

// MemberLoginResponse C 端登录成功响应。
type MemberLoginResponse struct {
	Member    mall.MallMember `json:"member"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}

// MemberProfileResponse C 端会员资料（含等级与积分）。
type MemberProfileResponse struct {
	Member  mall.MallMember        `json:"member"`
	Profile mall.MallMemberProfile `json:"profile"`
}
