// Package service 邮件插件业务层。
package service

// ServiceGroup 邮件 Service 分组。
type ServiceGroup struct {
	EmailService
}

var ServiceGroupApp = new(ServiceGroup)
