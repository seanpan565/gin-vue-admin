// mall_member_level.go 会员等级模型。
package mall

import "mall-admin/server/global"

// MallMemberLevel 会员等级配置表（如普通/银卡/金卡）。
type MallMemberLevel struct {
	global.GVA_MODEL
	Name         string `json:"name" gorm:"size:32;comment:等级名称"`
	Level        int    `json:"level" gorm:"uniqueIndex;comment:等级序号 数值越大等级越高"`
	MinGrowth    int64  `json:"minGrowth" gorm:"default:0;comment:升级所需成长值"`
	DiscountRate int    `json:"discountRate" gorm:"default:100;comment:折扣率 100=无折扣 95=95折"`
	Icon         string `json:"icon" gorm:"size:512;default:'';comment:等级图标"`
	Status       int    `json:"status" gorm:"default:1;comment:状态 1启用 2禁用"`
	Remark       string `json:"remark" gorm:"size:255;default:'';comment:备注"`
}

func (MallMemberLevel) TableName() string {
	return "mall_member_levels"
}
