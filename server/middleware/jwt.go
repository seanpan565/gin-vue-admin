// Package middleware Gin 中间件层，提供鉴权、权限、跨域、限流、日志等 HTTP 拦截能力。
// jwt.go JWT 鉴权中间件，校验 token 有效性并支持临近过期自动续签。
package middleware

import (
	"errors"
	"strconv"
	"time"

	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	"mall-admin/server/model/system"
	"mall-admin/server/service"
	"mall-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtUserService = service.ServiceGroupApp.SystemServiceGroup.UserService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

// JWTAuth 校验 x-token，处理黑名单、过期续签及多点登录状态。
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问，请登录", c)
			c.Abort()
			return
		}
		if isBlacklist(token) {
			response.NoAuth("您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth("登录已过期，请重新登录", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		if !isUserEnabled(claims.UUID.String()) {
			_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
			response.FailWithDetailed(gin.H{"reload": true}, "账号已禁用或不存在，请重新登录", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				// 记录新的活跃jwt
				_ = utils.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func isBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// isUserEnabled 校验用户是否可用，结果缓存 5 分钟以降低 DB 压力。
func isUserEnabled(uuid string) bool {
	cacheKey := utils.UserEnableCachePrefix + uuid
	if v, ok := global.BlackCache.Get(cacheKey); ok {
		enabled, _ := v.(bool)
		return enabled
	}
	user, err := jwtUserService.FindUserByUuid(uuid)
	enabled := err == nil && user.Enable != 2
	global.BlackCache.Set(cacheKey, enabled, 5*time.Minute)
	return enabled
}
