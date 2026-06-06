// mall_member_profile.go 会员扩展资料模型。
package mall

import (
	"time"

	"mall-admin/server/global"
)

// MallMemberProfile 会员扩展资料表，与 mall_members 一对一。
type MallMemberProfile struct {
	global.GVA_MODEL
	MemberID    uint            `json:"memberId" gorm:"uniqueIndex;comment:会员ID"`
	Gender      int8            `json:"gender" gorm:"default:0;comment:性别 0未知 1男 2女"`
	Birthday    *time.Time      `json:"birthday" gorm:"comment:生日"`
	LevelID     uint            `json:"levelId" gorm:"default:1;comment:会员等级ID"`
	Level       MallMemberLevel `json:"level" gorm:"foreignKey:LevelID;references:ID"`
	Points      int64           `json:"points" gorm:"default:0;comment:可用积分"`
	GrowthValue int64           `json:"growthValue" gorm:"default:0;comment:成长值"`
	TotalSpent  int64           `json:"totalSpent" gorm:"default:0;comment:累计消费金额(分)"`
	OrderCount  int             `json:"orderCount" gorm:"default:0;comment:累计订单数"`
}

func (MallMemberProfile) TableName() string {
	return "mall_member_profiles"
}
