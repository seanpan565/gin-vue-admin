package request

// 角色按钮权限相关请求参数
type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuID"`
	AuthorityId uint   `json:"authorityId"`
	Selected    []uint `json:"selected"`
}
