// Package model auto 插件数据模型。
package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// sys_auto_code_package.go 自动化代码包配置模型。

// SysAutoCodePackage 自动化代码包实体。
type SysAutoCodePackage struct {
	global.GVA_MODEL
	Desc        string `json:"desc" gorm:"comment:描述"`
	Label       string `json:"label" gorm:"comment:显示名称"`
	Template    string `json:"template" gorm:"comment:模板"`
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Module      string `json:"-" example:"模块"`
}

func (s *SysAutoCodePackage) TableName() string {
	return "sys_auto_code_packages"
}
