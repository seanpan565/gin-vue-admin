package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// JwtBlacklist JWT 黑名单表模型
type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
