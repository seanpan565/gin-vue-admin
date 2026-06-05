package response

import "github.com/flipped-aurora/gin-vue-admin/server/config"

// 系统配置相关响应结构
type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
