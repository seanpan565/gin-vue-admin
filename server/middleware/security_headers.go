// security_headers.go HTTP 安全响应头中间件。
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// SecurityHeaders 设置常见安全响应头，降低 XSS/点击劫持等风险。
func SecurityHeaders() gin.HandlerFunc {
	sec := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		ReferrerPolicy:     "strict-origin-when-cross-origin",
	})
	return func(c *gin.Context) {
		if err := sec.Process(c.Writer, c.Request); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}
