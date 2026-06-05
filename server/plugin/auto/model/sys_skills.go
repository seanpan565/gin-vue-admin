// Package model auto 插件数据模型。
package model

// sys_skills.go AI 技能元数据与详情结构。

// SkillMeta 技能元信息（YAML 头）。
type SkillMeta struct {
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
	AllowedTools string `json:"allowedTools" yaml:"allowed-tools,omitempty"`
	Context      string `json:"context" yaml:"context,omitempty"`
	Agent        string `json:"agent" yaml:"agent,omitempty"`
}

// SkillDetail 技能完整内容。
type SkillDetail struct {
	Tool       string    `json:"tool"`
	Skill      string    `json:"skill"`
	Meta       SkillMeta `json:"meta"`
	Markdown   string    `json:"markdown"`
	Scripts    []string  `json:"scripts"`
	Resources  []string  `json:"resources"`
	References []string  `json:"references"`
	Templates  []string  `json:"templates"`
}

// SkillTool 可用 AI 工具项。
type SkillTool struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

