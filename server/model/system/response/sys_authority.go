package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

// 角色管理相关响应结构
type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId"` // 旧角色ID
}
