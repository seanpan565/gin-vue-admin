// loadtls.go HTTPS 重定向中间件，配合 Router.RunTLS 使用。
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// LoadTls 强制将 HTTP 请求重定向到 HTTPS。
func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
