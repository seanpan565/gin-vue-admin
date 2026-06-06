// other.go 初始化 JWT 缓存等杂项配置。
package initialize

import (
	"mall-admin/server/global"
	"mall-admin/server/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

// OtherInit 解析 JWT 时长并初始化本地缓存。
func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
