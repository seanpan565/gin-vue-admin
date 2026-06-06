// Package request 公告插件请求参数模型。
package request

import (
	"mall-admin/server/model/common/request"
	"time"
)

// InfoSearch 公告分页搜索条件。
type InfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
