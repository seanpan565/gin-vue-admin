// member_token_blacklist.go C 端会员 Token 黑名单，复用 jwt_blacklists 表与 BlackCache。
package utils

import (
	"mall-admin/server/global"
	"mall-admin/server/model/system"
	"go.uber.org/zap"
)

// BlacklistMemberToken 将会员 token 写入黑名单（持久化 + 内存缓存）。
func BlacklistMemberToken(token string) {
	if token == "" {
		return
	}
	record := system.JwtBlacklist{Jwt: token}
	if err := global.GVA_DB.Create(&record).Error; err != nil {
		global.GVA_LOG.Error("会员 token 拉黑失败", zap.Error(err))
		return
	}
	global.BlackCache.SetDefault(token, struct{}{})
}

// IsMemberTokenBlacklisted 判断会员 token 是否已失效。
func IsMemberTokenBlacklisted(token string) bool {
	if token == "" {
		return false
	}
	_, ok := global.BlackCache.Get(token)
	return ok
}
