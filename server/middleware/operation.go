// operation.go 操作记录中间件，将请求/响应信息写入 sys_operation_records 表。
package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"mall-admin/server/utils"

	"mall-admin/server/global"
	"mall-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

// OperationRecord 记录用户 API 操作详情（请求体、响应、耗时等）。
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.GVA_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.GetClaims(c)
		if claims != nil && claims.BaseClaims.ID != 0 {
			userId = int(claims.BaseClaims.ID)
		} else {
			id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   "",
			UserID: userId,
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			record.Body = "[文件]"
		} else {
			safeBody := desensitizeRequestBody(c.Request.URL.Path, body)
			if len(safeBody) > bufferSize {
				record.Body = "[超出记录长度]"
			} else {
				record.Body = string(safeBody)
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > bufferSize {
				// 截断
				record.Body = "超出记录长度"
			}
		}
		if err := global.GVA_DB.Create(&record).Error; err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

var sensitiveBodyPaths = []string{
	"/base/login",
	"/user/changePassword",
	"/user/resetPassword",
	"/user/admin_register",
	"/site/auth/login",
	"/site/auth/register",
}

var sensitiveBodyFields = []string{
	"password", "passWord", "newPassword", "oldPassword",
}

// desensitizeRequestBody 对敏感接口请求体中的密码字段脱敏。
func desensitizeRequestBody(path string, body []byte) []byte {
	if len(body) == 0 {
		return body
	}
	needMask := false
	for _, p := range sensitiveBodyPaths {
		if strings.HasSuffix(path, p) {
			needMask = true
			break
		}
	}
	if !needMask {
		return body
	}
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return []byte("[敏感请求已脱敏]")
	}
	for _, field := range sensitiveBodyFields {
		if _, ok := payload[field]; ok {
			payload[field] = "***"
		}
	}
	masked, err := json.Marshal(payload)
	if err != nil {
		return []byte("[敏感请求已脱敏]")
	}
	return masked
}
