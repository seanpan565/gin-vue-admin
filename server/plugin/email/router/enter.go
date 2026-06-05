// Package router 邮件插件路由层。
package router

// RouterGroup 邮件路由分组。
type RouterGroup struct {
	EmailRouter
}

var RouterGroupApp = new(RouterGroup)
