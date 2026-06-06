package response

import (
	"mall-admin/server/model/system/request"
)

// Casbin 权限相关响应结构
type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
