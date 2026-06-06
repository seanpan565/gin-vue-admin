// error.go Panic 恢复中间件，记录堆栈并将错误写入 sys_errors 表。
package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"mall-admin/server/global"
	"mall-admin/server/model/system"
	"mall-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinRecovery 捕获 panic 并记录日志，可选将堆栈写入 sys_errors 表。
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.GVA_LOG.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					form := "后端"
					info := fmt.Sprintf("Panic: %v\nRequest: %s\nStack: %s", err, string(httpRequest), string(debug.Stack()))
					level := "error"
					_ = service.ServiceGroupApp.SystemServiceGroup.SysErrorService.CreateSysError(context.Background(), &system.SysError{
						Form:  &form,
						Info:  &info,
						Level: level,
					})
					global.GVA_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				} else {
					form := "后端"
					info := fmt.Sprintf("Panic: %v\nRequest: %s", err, string(httpRequest))
					level := "error"
					_ = service.ServiceGroupApp.SystemServiceGroup.SysErrorService.CreateSysError(context.Background(), &system.SysError{
						Form:  &form,
						Info:  &info,
						Level: level,
					})
					global.GVA_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
