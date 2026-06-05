// model.go 提供 GORM 基础模型，包含主键、时间戳与软删除字段。
package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
