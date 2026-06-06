// Package router 邮件插件路由层。
package router

// sys_email.go 邮件 HTTP 路由注册。

import (
	"mall-admin/server/middleware"
	"mall-admin/server/plugin/email/api"
	"github.com/gin-gonic/gin"
)

// EmailRouter 邮件路由组。
type EmailRouter struct{}

// InitEmailRouter 注册邮件发送路由。
func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
