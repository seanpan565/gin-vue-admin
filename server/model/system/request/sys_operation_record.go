package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 操作记录相关请求参数
type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
