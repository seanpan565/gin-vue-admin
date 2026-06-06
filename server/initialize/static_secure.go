// static_secure.go 上传目录安全静态文件服务。
package initialize

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"mall-admin/server/global"
	"github.com/gin-gonic/gin"
)

var blockedStaticExts = map[string]struct{}{
	".php": {}, ".jsp": {}, ".asp": {}, ".aspx": {}, ".exe": {},
	".sh": {}, ".bat": {}, ".cmd": {}, ".ps1": {}, ".html": {}, ".htm": {},
	".js": {}, ".mjs": {}, ".cjs": {}, ".cgi": {}, ".pl": {},
}

// RegisterSecureStatic 注册带安全头的上传文件静态路由。
func RegisterSecureStatic(router *gin.Engine) {
	prefix := "/" + strings.Trim(global.GVA_CONFIG.Local.StorePath, "/")
	router.GET(prefix+"/*filepath", serveSecureStatic)
}

func serveSecureStatic(c *gin.Context) {
	rel := strings.TrimPrefix(c.Param("filepath"), "/")
	ext := strings.ToLower(filepath.Ext(rel))
	if _, blocked := blockedStaticExts[ext]; blocked {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	base := global.GVA_CONFIG.Local.StorePath
	full := filepath.Join(base, rel)
	if !strings.HasPrefix(filepath.Clean(full), filepath.Clean(base)) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if _, err := os.Stat(full); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Content-Security-Policy", "default-src 'none'")
	c.File(full)
}
