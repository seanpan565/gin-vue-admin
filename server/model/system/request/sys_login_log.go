package request

import (
	"mall-admin/server/model/common/request"
	"mall-admin/server/model/system"
)

// 登录日志相关请求参数
type SysLoginLogSearch struct {
	system.SysLoginLog
	request.PageInfo
}
