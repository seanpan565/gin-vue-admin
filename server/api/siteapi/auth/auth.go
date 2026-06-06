package auth

import (
	"mall-admin/server/model/common/response"
	"mall-admin/server/model/mall"
	mallReq "mall-admin/server/model/mall/request"
	mallRes "mall-admin/server/model/mall/response"
	"mall-admin/server/utils"
	"github.com/gin-gonic/gin"
)

// AuthApi C 端会员认证接口。
type AuthApi struct{}

// Register
// @Tags     MallSiteAuth
// @Summary  会员注册
// @Produce  application/json
// @Param    data  body      mallReq.MemberRegister  true  "手机号, 密码, 昵称"
// @Success  200   {object}  response.Response{data=mallRes.MemberLoginResponse,msg=string}
// @Router   /site/auth/register [post]
func (a *AuthApi) Register(c *gin.Context) {
	var req mallReq.MemberRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req, utils.MemberRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !utils.VerifyMemberCaptcha(c, req.CaptchaId, req.Captcha) {
		respondMemberLogin(c, mall.MallMember{}, "", false, req.Mobile, "验证码错误")
		return
	}

	member, err := memberService.Register(req, c.ClientIP())
	if err != nil {
		respondMemberLogin(c, member, "", false, req.Mobile, err.Error())
		return
	}
	respondMemberLogin(c, member, "注册成功", true, req.Mobile, "")
}

// Login
// @Tags     MallSiteAuth
// @Summary  会员登录
// @Produce  application/json
// @Param    data  body      mallReq.MemberLogin  true  "手机号, 密码"
// @Success  200   {object}  response.Response{data=mallRes.MemberLoginResponse,msg=string}
// @Router   /site/auth/login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var req mallReq.MemberLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req, utils.MemberLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !utils.VerifyMemberCaptcha(c, req.CaptchaId, req.Captcha) {
		respondMemberLogin(c, mall.MallMember{}, "", false, req.Mobile, "验证码错误")
		return
	}

	member, err := memberService.Login(req.Mobile, req.Password)
	if err != nil {
		respondMemberLogin(c, mall.MallMember{}, "", false, req.Mobile, err.Error())
		return
	}
	respondMemberLogin(c, *member, "登录成功", true, req.Mobile, "")
}

// Profile
// @Tags     MallSiteAuth
// @Summary  获取当前会员资料
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200  {object}  response.Response{data=mallRes.MemberProfileResponse,msg=string}
// @Router   /site/auth/profile [get]
func (a *AuthApi) Profile(c *gin.Context) {
	memberID := utils.GetMemberID(c)
	if memberID == 0 {
		response.NoAuth("请先登录", c)
		return
	}

	member, err := memberService.GetMemberByID(memberID)
	if err != nil {
		response.FailWithMessage("会员不存在", c)
		return
	}
	member.Password = ""

	response.OkWithDetailed(mallRes.MemberProfileResponse{
		Member:  member,
		Profile: member.Profile,
	}, "获取成功", c)
}

// Logout
// @Tags     MallSiteAuth
// @Summary  会员退出登录
// @Produce  application/json
// @Success  200  {object}  response.Response{msg=string}
// @Router   /site/auth/logout [post]
func (a *AuthApi) Logout(c *gin.Context) {
	if token := utils.GetMemberToken(c); token != "" {
		utils.BlacklistMemberToken(token)
	}
	utils.ClearMemberToken(c)
	response.OkWithMessage("已退出登录", c)
}
