package request

import (
	"mall-admin/server/model/common/request"
	"mall-admin/server/model/system"
)

// 操作记录相关请求参数
type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
