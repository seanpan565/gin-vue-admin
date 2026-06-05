// Package response 定义系统模块 API 响应结构。
package response

import "time"

// AI 工作流会话相关响应结构
type SysAIWorkflowSessionListItem struct {
	ID             uint      `json:"ID"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      time.Time `json:"UpdatedAt"`
	Tab            string    `json:"tab"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	ConversationID string    `json:"conversationId"`
	CurrentNodeID  string    `json:"currentNodeId"`
}

type AIWorkflowMarkdownDumpResult struct {
	FileName     string `json:"fileName"`
	FilePath     string `json:"filePath"`
	RelativePath string `json:"relativePath"`
	Directory    string `json:"directory"`
}
