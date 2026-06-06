// member_jwt.go C 端会员 JWT 鉴权中间件。
package middleware

import (
	"errors"
	"strconv"
	"time"

	"mall-admin/server/global"
	mallReq "mall-admin/server/model/mall/request"
	"mall-admin/server/model/common/response"
	"mall-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// MemberJWT C 端会员鉴权中间件，读取请求头 x-member-token，不走 Casbin。
func MemberJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetMemberToken(c)
		if token == "" {
			response.NoAuth("请先登录", c)
			c.Abort()
			return
		}
		if utils.IsMemberTokenBlacklisted(token) {
			response.NoAuth("登录已失效，请重新登录", c)
			utils.ClearMemberToken(c)
			c.Abort()
			return
		}

		j := utils.NewMemberJWT()
		claims, err := j.ParseMemberToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth("登录已过期，请重新登录", c)
			} else {
				response.NoAuth("无效的登录状态", c)
			}
			utils.ClearMemberToken(c)
			c.Abort()
			return
		}

		c.Set("memberClaims", claims)

		if claims.ExpiresAt != nil && claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, err := j.CreateMemberToken(*claims)
			if err == nil {
				newClaims, _ := j.ParseMemberToken(newToken)
				utils.SetMemberToken(c, newToken)
				if newClaims != nil && newClaims.ExpiresAt != nil {
					c.Header("x-member-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
				}
				c.Set("memberClaims", newClaims)
			}
		}

		c.Next()
	}
}

// GetMemberClaimsFromContext 从 gin.Context 读取会员 Claims。
func GetMemberClaimsFromContext(c *gin.Context) (*mallReq.MemberCustomClaims, bool) {
	v, ok := c.Get("memberClaims")
	if !ok {
		return nil, false
	}
	claims, ok := v.(*mallReq.MemberCustomClaims)
	return claims, ok
}
