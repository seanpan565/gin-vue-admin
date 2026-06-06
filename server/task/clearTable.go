// Package task 定时任务，维护数据库表数据。
package task

// clearTable.go 按保留周期清理过期表记录。

import (
	"errors"
	"fmt"
	"mall-admin/server/model/common"
	"time"

	"gorm.io/gorm"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: ClearTable
//@description: 清理数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

// ClearTable 清理操作记录与 JWT 黑名单等过期数据。
func ClearTable(db *gorm.DB) error {
	var ClearTableDetail []common.ClearDB

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "mall_member_login_logs",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	if db == nil {
		return errors.New("db Cannot be empty")
	}

	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		err = db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}
