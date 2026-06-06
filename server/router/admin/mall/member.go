// member.go 商城后台会员路由注册。
package mall

import "github.com/gin-gonic/gin"

// MemberRouter 后台会员管理路由。
type MemberRouter struct{}

// InitMemberRouter 注册 /mall/member/* 路由（需后台 JWT + Casbin）。
func (r *MemberRouter) InitMemberRouter(Router *gin.RouterGroup) {
	member := Router.Group("mall").Group("member")
	{
		member.POST("list", memberApi.GetMemberList)
		member.POST("detail", memberApi.GetMemberDetail)
	}
}
