// Package main 公告 GORM Gen 代码生成入口。
package main

// gen.go 生成公告 DAO 层代码。
//go:generate go mod tidy
//go:generate go mod download
//go:generate go run gen.go

import (
	"gorm.io/gen"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
)

// main 执行 GORM Gen 生成。
func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "announcement", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(
		new(model.Info),
	)
	g.Execute()
}
