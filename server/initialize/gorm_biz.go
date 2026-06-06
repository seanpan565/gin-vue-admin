// gorm_biz.go 业务扩展表的自动迁移入口，供 RegisterTables 调用。
package initialize

import (
	"mall-admin/server/global"
	"mall-admin/server/model/mall"
	mallService "mall-admin/server/service/mall"
)

// bizModel 迁移业务扩展表，开发者在此追加自定义 model。
func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		&mall.MallMemberLevel{},
		&mall.MallMember{},
		&mall.MallMemberProfile{},
		&mall.MallMemberOauth{},
		&mall.MallMemberLoginLog{},
	)
	if err != nil {
		return err
	}
	return seedMallDefaults()
}

// seedMallDefaults 写入商城模块默认数据。
func seedMallDefaults() error {
	return new(mallService.MemberService).EnsureDefaultLevels()
}
