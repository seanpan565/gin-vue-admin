// Package initialize 公告插件安装初始化。
package initialize

// gorm.go 安装时迁移公告数据表。

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Gorm 自动迁移公告 Info 表。
func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(
		new(model.Info),
	)
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
