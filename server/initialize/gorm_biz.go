// gorm_biz.go 业务扩展表的自动迁移入口，供 RegisterTables 调用。
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// bizModel 迁移业务扩展表，开发者在此追加自定义 model。
func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
