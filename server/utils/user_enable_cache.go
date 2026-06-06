// user_enable_cache.go 管理端用户启用状态缓存（JWT 中间件使用）。
package utils

import "mall-admin/server/global"

const UserEnableCachePrefix = "user_enable:"

// InvalidateUserEnableCache 用户信息变更（尤其 enable）后清除缓存。
func InvalidateUserEnableCache(uuid string) {
	if uuid == "" {
		return
	}
	global.BlackCache.Delete(UserEnableCachePrefix + uuid)
}
