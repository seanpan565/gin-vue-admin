// Package mcpTool MCP 请求上下文与鉴权令牌传递。
package mcpTool

import (
	"context"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type mcpContextKey string

const authTokenContextKey mcpContextKey = "mcp-auth-token"

// WithHTTPRequestContext 将请求鉴权 token 写入 context。
func WithHTTPRequestContext(ctx context.Context, r *http.Request) context.Context {
	token := extractIncomingAuthToken(r.Header)
	return context.WithValue(ctx, authTokenContextKey, token)
}

func configuredAuthHeader() string {
	if header := strings.TrimSpace(global.GVA_CONFIG.MCP.AuthHeader); header != "" {
		return header
	}
	return "x-token"
}

// ConfiguredAuthHeader 返回配置的鉴权请求头名称。
func ConfiguredAuthHeader() string {
	return configuredAuthHeader()
}

func authTokenFromContext(ctx context.Context) string {
	token, _ := ctx.Value(authTokenContextKey).(string)
	return strings.TrimSpace(token)
}

func extractIncomingAuthToken(headers http.Header) string {
	candidates := []string{
		configuredAuthHeader(),
		"x-token",
		"token",
		"authorization",
	}

	seen := make(map[string]struct{}, len(candidates))
	for _, name := range candidates {
		key := strings.ToLower(strings.TrimSpace(name))
		if key == "" {
			continue
		}
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}

		value := strings.TrimSpace(headers.Get(name))
		if value == "" {
			continue
		}
		if key == "authorization" {
			return strings.TrimSpace(strings.TrimPrefix(value, "Bearer "))
		}
		return value
	}

	return ""
}
