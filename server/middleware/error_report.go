// error_report.go 前端错误上报接口防护。
package middleware

import (
	"net/http"

	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

// ErrorReportGuard 校验错误上报密钥（配置为空时仅依赖限流）。
func ErrorReportGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := global.GVA_CONFIG.Security.ErrorReportKey
		if key != "" && c.GetHeader("X-Error-Report-Key") != key {
			response.FailWithMessage("无权上报错误", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

// MetricsAuth Prometheus 指标端点鉴权。
func MetricsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := global.GVA_CONFIG.Security.MetricsAuthToken
		if token == "" {
			c.Next()
			return
		}
		if c.GetHeader("Authorization") != "Bearer "+token && c.Query("token") != token {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
