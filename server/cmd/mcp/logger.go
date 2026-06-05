// MCP 独立服务日志初始化（开发模式 zap）
package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func initializeStandaloneLogger() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	global.GVA_LOG = logger
	zap.ReplaceGlobals(logger)
	return nil
}
