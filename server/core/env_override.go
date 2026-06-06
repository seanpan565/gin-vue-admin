// env_override.go 用环境变量覆盖敏感配置，避免密钥写入仓库。
package core

import (
	"os"
	"strings"

	"mall-admin/server/global"
)

// ApplyEnvOverrides 在 Viper 加载后应用环境变量覆盖。
func ApplyEnvOverrides() {
	if v := os.Getenv("GVA_JWT_SIGNING_KEY"); v != "" {
		global.GVA_CONFIG.JWT.SigningKey = v
	}
	if v := os.Getenv("GVA_MYSQL_PASSWORD"); v != "" {
		global.GVA_CONFIG.Mysql.Password = v
	}
	if v := os.Getenv("GVA_REDIS_PASSWORD"); v != "" {
		global.GVA_CONFIG.Redis.Password = v
	}
	if v := os.Getenv("GVA_SECURITY_ERROR_REPORT_KEY"); v != "" {
		global.GVA_CONFIG.Security.ErrorReportKey = v
	}
	if v := os.Getenv("GVA_SECURITY_METRICS_AUTH_TOKEN"); v != "" {
		global.GVA_CONFIG.Security.MetricsAuthToken = v
	}
	if v := strings.ToLower(os.Getenv("GVA_SECURITY_SWAGGER_ENABLE")); v == "true" || v == "1" {
		global.GVA_CONFIG.Security.SwaggerEnable = true
	} else if v == "false" || v == "0" {
		global.GVA_CONFIG.Security.SwaggerEnable = false
	}
}
