// mall_member_oauth.go 会员第三方登录绑定模型。
package mall

import "mall-admin/server/global"

// MallMemberOauth 第三方账号绑定表（微信/小程序等），预留扩展。
type MallMemberOauth struct {
	global.GVA_MODEL
	MemberID uint   `json:"memberId" gorm:"index;comment:会员ID"`
	Provider string `json:"provider" gorm:"size:32;uniqueIndex:idx_mall_oauth_provider_openid,priority:1;comment:平台 wechat/miniprogram/app"`
	OpenID   string `json:"openId" gorm:"size:128;uniqueIndex:idx_mall_oauth_provider_openid,priority:2;comment:平台OpenID"`
	UnionID  string `json:"unionId" gorm:"size:128;index;default:'';comment:UnionID"`
	Extra    string `json:"extra" gorm:"type:text;comment:扩展信息JSON"`
}

func (MallMemberOauth) TableName() string {
	return "mall_member_oauths"
}
