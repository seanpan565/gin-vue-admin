package auth

import (
	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	"mall-admin/server/model/mall"
	mallReq "mall-admin/server/model/mall/request"
	mallRes "mall-admin/server/model/mall/response"
	"mall-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func respondMemberLogin(c *gin.Context, member mall.MallMember, msg string, loginOK bool, logMobile string, logErr string) {
	if !loginOK {
		_ = memberService.CreateLoginLog(mall.MallMemberLoginLog{
			Mobile:       logMobile,
			IP:           c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			Status:       false,
			ErrorMessage: logErr,
		})
		response.FailWithMessage(logErr, c)
		return
	}

	token, claims, err := utils.MemberLoginToken(mallReq.MemberBaseClaims{
		UUID:     member.UUID,
		ID:       member.ID,
		Mobile:   member.Mobile,
		Nickname: member.Nickname,
	})
	if err != nil {
		global.GVA_LOG.Error("会员 token 生成失败", zap.Error(err))
		response.FailWithMessage("登录失败，请稍后重试", c)
		return
	}

	if err := memberService.UpdateLastLogin(member.ID, c.ClientIP()); err != nil {
		global.GVA_LOG.Warn("更新会员登录时间失败", zap.Error(err))
	}
	if err := memberService.CreateLoginLog(mall.MallMemberLoginLog{
		MemberID:  member.ID,
		Mobile:    member.Mobile,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Status:    true,
	}); err != nil {
		global.GVA_LOG.Warn("写入会员登录日志失败", zap.Error(err))
	}

	utils.SetMemberToken(c, token)
	member.Password = ""
	response.OkWithDetailed(mallRes.MemberLoginResponse{
		Member:    member,
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}, msg, c)
}
