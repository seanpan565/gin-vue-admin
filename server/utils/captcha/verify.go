// verify.go 验证码校验封装。
package captcha

import (
	"context"
	"time"

	"mall-admin/server/global"
)

// VerifyCaptcha 校验图形验证码，失败时递增 IP 计数。
func VerifyCaptcha(ctx context.Context, clientIP, captchaID, captcha string) bool {
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut
	v, ok := global.BlackCache.Get(clientIP)
	if !ok {
		global.BlackCache.Set(clientIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	needCaptcha := openCaptcha == 0 || openCaptcha < toInt(v)
	if !needCaptcha {
		return true
	}
	if captchaID == "" || captcha == "" {
		return false
	}
	if !Store(ctx).Verify(captchaID, captcha, true) {
		global.BlackCache.Increment(clientIP, 1)
		return false
	}
	return true
}

func toInt(v interface{}) int {
	switch n := v.(type) {
	case int:
		return n
	case int64:
		return int(n)
	default:
		return 0
	}
}
