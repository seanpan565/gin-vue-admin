// member_auth.go C 端会员注册/登录请求参数。
package request

// MemberRegister 会员注册请求。
type MemberRegister struct {
	Mobile         string `json:"mobile" binding:"required"`
	Password       string `json:"password" binding:"required,min=6,max=32"`
	Nickname       string `json:"nickname"`
	RegisterSource string `json:"registerSource"` // h5 / app / miniprogram，默认 h5
	Captcha        string `json:"captcha"`
	CaptchaId      string `json:"captchaId"`
}

// MemberLogin 会员登录请求。
type MemberLogin struct {
	Mobile    string `json:"mobile" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
