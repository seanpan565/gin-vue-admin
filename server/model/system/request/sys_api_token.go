package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// API 令牌相关请求参数
type SysApiTokenSearch struct {
	system.SysApiToken
	request.PageInfo
    Status *bool `json:"status" form:"status"`
}
