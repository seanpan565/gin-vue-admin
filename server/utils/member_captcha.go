// member_captcha.go C 端会员验证码校验。
package utils

import (
	"mall-admin/server/global"
	captchaStore "mall-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
)

// VerifyMemberCaptcha 按配置校验会员端验证码。
func VerifyMemberCaptcha(c *gin.Context, captchaID, captcha string) bool {
	if !global.GVA_CONFIG.Security.MemberCaptchaEnable {
		return true
	}
	return captchaStore.VerifyCaptcha(c.Request.Context(), c.ClientIP(), captchaID, captcha)
}
