// Package global 邮件插件全局变量。
package global

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/config"

// GlobalConfig 邮件插件全局 SMTP 配置。
var GlobalConfig = new(config.Email)
