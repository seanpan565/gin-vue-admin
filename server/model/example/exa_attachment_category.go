// Package example 定义示例模块数据库实体模型。
package example

import (
	"mall-admin/server/global"
)

// ExaAttachmentCategory 附件分类表模型
type ExaAttachmentCategory struct {
	global.GVA_MODEL
	Name     string                   `json:"name" form:"name" gorm:"default:null;type:varchar(255);column:name;comment:分类名称;"`
	Pid      uint                     `json:"pid" form:"pid" gorm:"default:0;type:int;column:pid;comment:父节点ID;"`
	Children []*ExaAttachmentCategory `json:"children" gorm:"-"`
}

func (ExaAttachmentCategory) TableName() string {
	return "exa_attachment_category"
}
