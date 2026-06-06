// member_auth.go C 端会员认证路由注册。
package siteapi

import (
	"mall-admin/server/middleware"

	"github.com/gin-gonic/gin"
)

// MemberAuthRouter C 端会员认证路由。
type MemberAuthRouter struct{}

// InitMemberAuthRouter 注册 /site/auth/* 路由。
func (r *MemberAuthRouter) InitMemberAuthRouter(Router *gin.RouterGroup) {
	auth := Router.Group("auth").Use(middleware.AuthRateLimit())
	{
		auth.POST("register", authApi.Register)
		auth.POST("login", authApi.Login)
		auth.POST("logout", authApi.Logout)
	}

	authed := Router.Group("auth").Use(middleware.MemberJWT())
	{
		authed.GET("profile", authApi.Profile)
	}
}
