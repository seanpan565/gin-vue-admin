// member_token_blacklist.go C 端会员 Token 黑名单（内存缓存，登出后失效）。
package utils

import "mall-admin/server/global"

const memberTokenBlacklistPrefix = "member_jwt:"

// BlacklistMemberToken 将会员 token 加入黑名单。
func BlacklistMemberToken(token string) {
	if token == "" {
		return
	}
	global.BlackCache.SetDefault(memberTokenBlacklistPrefix+token, struct{}{})
}

// IsMemberTokenBlacklisted 判断会员 token 是否已登出失效。
func IsMemberTokenBlacklisted(token string) bool {
	if token == "" {
		return false
	}
	_, ok := global.BlackCache.Get(memberTokenBlacklistPrefix + token)
	return ok
}
