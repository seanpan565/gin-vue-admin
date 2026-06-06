package response

import "mall-admin/server/config"

// 系统配置相关响应结构
type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
