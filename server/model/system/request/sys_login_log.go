package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 登录日志相关请求参数
type SysLoginLogSearch struct {
	system.SysLoginLog
	request.PageInfo
}
