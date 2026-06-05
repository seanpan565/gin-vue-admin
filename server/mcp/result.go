// Package mcpTool MCP 工具返回结果格式化。
package mcpTool

// result.go 将结构化数据序列化为 MCP 文本结果。

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// textResultWithJSON 生成带 JSON 内容的 MCP 文本结果。
func textResultWithJSON(title string, payload any) (*mcp.CallToolResult, error) {
	resultJSON, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("序列化结果失败: %w", err)
	}

	text := string(resultJSON)
	if title != "" {
		text = fmt.Sprintf("%s\n\n%s", title, text)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
	}, nil
}
