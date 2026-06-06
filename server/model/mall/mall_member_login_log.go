// mall_member_login_log.go 会员登录日志模型。
package mall

import "mall-admin/server/global"

// MallMemberLoginLog 会员登录日志表。
type MallMemberLoginLog struct {
	global.GVA_MODEL
	MemberID     uint   `json:"memberId" gorm:"index;comment:会员ID"`
	Mobile       string `json:"mobile" gorm:"size:20;comment:手机号"`
	IP           string `json:"ip" gorm:"size:64;comment:登录IP"`
	UserAgent    string `json:"userAgent" gorm:"size:512;comment:UserAgent"`
	Status       bool   `json:"status" gorm:"comment:是否成功"`
	ErrorMessage string `json:"errorMessage" gorm:"size:255;default:'';comment:失败原因"`
}

func (MallMemberLoginLog) TableName() string {
	return "mall_member_login_logs"
}
