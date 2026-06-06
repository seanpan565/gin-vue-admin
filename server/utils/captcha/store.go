package captcha

import (
	"context"

	"mall-admin/server/global"
	"github.com/mojocn/base64Captcha"
)

// Store 按配置返回验证码存储：启用 Redis 时用 Redis，否则用进程内存。
func Store(ctx context.Context) base64Captcha.Store {
	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		return NewDefaultRedisStore().UseWithCtx(ctx)
	}
	return base64Captcha.DefaultMemStore
}
