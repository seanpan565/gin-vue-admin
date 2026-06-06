package system

import (
	"mall-admin/server/config"
)

// System 系统运行时配置视图模型
// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
