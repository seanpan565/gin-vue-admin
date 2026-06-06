// health.go 健康检查与就绪探针。
package initialize

import (
	"context"
	"net/http"
	"time"

	"mall-admin/server/global"
	"github.com/gin-gonic/gin"
)

// HealthHandler 返回服务健康状态，供负载均衡与监控使用。
func HealthHandler(c *gin.Context) {
	status := gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	}
	code := http.StatusOK

	if global.GVA_DB != nil {
		sqlDB, err := global.GVA_DB.DB()
		if err != nil || sqlDB.PingContext(c.Request.Context()) != nil {
			status["database"] = "down"
			code = http.StatusServiceUnavailable
		} else {
			status["database"] = "up"
		}
	} else {
		status["database"] = "not_initialized"
	}

	if global.GVA_CONFIG.System.UseRedis {
		if global.GVA_REDIS == nil {
			status["redis"] = "down"
			code = http.StatusServiceUnavailable
		} else if err := global.GVA_REDIS.Ping(context.Background()).Err(); err != nil {
			status["redis"] = "down"
			code = http.StatusServiceUnavailable
		} else {
			status["redis"] = "up"
		}
	}

	c.JSON(code, status)
}
