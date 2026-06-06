// member_admin.go 后台会员管理请求参数。
package request

import (
	"mall-admin/server/model/common/request"
)

// MemberSearch 后台会员列表查询。
type MemberSearch struct {
	request.PageInfo
	Mobile string `json:"mobile" form:"mobile"`
	Status *int   `json:"status" form:"status"`
}

// MemberByID 按 ID 查询会员。
type MemberByID struct {
	ID uint `json:"id" binding:"required"`
}
