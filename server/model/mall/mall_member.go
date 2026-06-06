// mall_member.go 商城 C 端会员账号模型。
package mall

import (
	"time"

	"mall-admin/server/global"
	"github.com/google/uuid"
)

// MallMember 商城 C 端会员账号表，与 sys_users 完全独立。
type MallMember struct {
	global.GVA_MODEL
	UUID           uuid.UUID         `json:"uuid" gorm:"index;comment:会员UUID"`
	Mobile         string            `json:"mobile" gorm:"uniqueIndex;size:20;comment:手机号"`
	Email          string            `json:"email" gorm:"index;size:128;default:'';comment:邮箱"`
	Password       string            `json:"-" gorm:"comment:登录密码(bcrypt)"`
	Nickname       string            `json:"nickname" gorm:"size:64;comment:昵称"`
	Avatar         string            `json:"avatar" gorm:"size:512;default:'';comment:头像URL"`
	Status         int               `json:"status" gorm:"default:1;comment:状态 1正常 2冻结"`
	RegisterSource string            `json:"registerSource" gorm:"size:32;default:h5;comment:注册来源 h5/app/wechat/miniprogram"`
	RegisterIP     string            `json:"registerIp" gorm:"size:64;default:'';comment:注册IP"`
	LastLoginAt    *time.Time        `json:"lastLoginAt" gorm:"comment:最后登录时间"`
	LastLoginIP    string            `json:"lastLoginIp" gorm:"size:64;default:'';comment:最后登录IP"`
	Profile        MallMemberProfile `json:"profile" gorm:"foreignKey:MemberID;references:ID"`
}

func (MallMember) TableName() string {
	return "mall_members"
}
