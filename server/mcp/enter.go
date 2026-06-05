// Package mcpTool MCP 工具注册中心，统一管理 AI 可调用工具。
package mcpTool

// enter.go 定义 MCP 工具接口与注册表。

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// McpTool 定义了MCP工具必须实现的接口
// McpTool MCP 工具必须实现的接口。
type McpTool interface {
	// Handle 返回工具调用信息
	Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
	// New 返回工具注册信息
	New() mcp.Tool
}

// 工具注册表
var toolRegister = make(map[string]McpTool)

// RegisterTool 供工具在init时调用，将自己注册到工具注册表中
// RegisterTool 工具 init 时注册到全局表。
func RegisterTool(tool McpTool) {
	mcpTool := tool.New()
	toolRegister[mcpTool.Name] = tool
}

// RegisterAllTools 将所有注册的工具注册到MCP服务中
// RegisterAllTools 将全部工具挂载到 MCP 服务。
func RegisterAllTools(mcpServer *server.MCPServer) {
	for _, tool := range toolRegister {
		mcpServer.AddTool(tool.New(), tool.Handle)
	}
}
