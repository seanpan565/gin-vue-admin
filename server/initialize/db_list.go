// db_list.go 初始化多数据库连接列表，支持别名切换主库。
package initialize

import (
	"mall-admin/server/config"
	"mall-admin/server/global"
	"gorm.io/gorm"
)

const sys = "system"

// DBList 遍历配置初始化多个数据库连接，alias 为 system 时覆盖主库。
func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GVA_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":
			dbMap[info.AliasName] = GormMssqlByConfig(config.Mssql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		case "oracle":
			dbMap[info.AliasName] = GormOracleByConfig(config.Oracle{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.GVA_DB = sysDB
	}
	global.GVA_DBList = dbMap
}
