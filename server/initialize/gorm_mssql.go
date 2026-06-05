// gorm_mssql.go SQL Server 数据库连接初始化。
package initialize

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql 初始化Mssql数据库
// Author [LouisZhang](191180776@qq.com)
func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mssql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(general)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
		return db
	}
}

// GormMssqlByConfig 初始化Mysql数据库用过传入配置
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
		return db
	}
}
