// member_claims.go C 端会员 JWT Claims 定义。
package utils

import (
	"mall-admin/server/global"
	mallReq "mall-admin/server/model/mall/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const memberTokenHeader = "x-member-token"

// GetMemberToken 从请求头获取 C 端会员 token。
func GetMemberToken(c *gin.Context) string {
	return c.GetHeader(memberTokenHeader)
}

// GetMemberClaims 解析 C 端会员 JWT。
func GetMemberClaims(c *gin.Context) (*mallReq.MemberCustomClaims, error) {
	token := GetMemberToken(c)
	if token == "" {
		return nil, TokenInvalid
	}
	return NewMemberJWT().ParseMemberToken(token)
}

// GetMemberID 从 Context 或 token 中获取会员 ID。
func GetMemberID(c *gin.Context) uint {
	if claims, exists := c.Get("memberClaims"); exists {
		return claims.(*mallReq.MemberCustomClaims).MemberBaseClaims.ID
	}
	if cl, err := GetMemberClaims(c); err == nil {
		return cl.MemberBaseClaims.ID
	}
	return 0
}

// GetMemberUUID 从 Context 或 token 中获取会员 UUID。
func GetMemberUUID(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("memberClaims"); exists {
		return claims.(*mallReq.MemberCustomClaims).MemberBaseClaims.UUID
	}
	if cl, err := GetMemberClaims(c); err == nil {
		return cl.MemberBaseClaims.UUID
	}
	return uuid.UUID{}
}

// SetMemberToken 将会员 token 写入响应头（供前端读取）。
func SetMemberToken(c *gin.Context, token string) {
	c.Header(memberTokenHeader, token)
}

// ClearMemberToken 清除响应头中的会员 token。
func ClearMemberToken(c *gin.Context) {
	c.Header(memberTokenHeader, "")
}

// LogMemberTokenError 记录会员 token 解析失败日志。
func LogMemberTokenError(msg string, err error) {
	if err != nil {
		global.GVA_LOG.Sugar().Debugw(msg, "error", err.Error())
	}
}
