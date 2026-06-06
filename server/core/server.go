// Package core 服务端核心层，负责启动 HTTP 服务及配置/日志等基础设施初始化。
package core

import (
	"fmt"
	"time"

	"mall-admin/server/global"
	"mall-admin/server/initialize"
	"mall-admin/server/service/system"
	"go.uber.org/zap"
)

// RunServer 初始化缓存与数据层，注册路由并启动 HTTP 服务。
func RunServer() {
	if global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	if global.GVA_CONFIG.System.UseMongo {
		if err := initialize.Mongo.Initialization(); err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}

	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	fmt.Printf(`
	商城管理后台
	当前版本:%s
	Swagger 文档:http://127.0.0.1%s/swagger/index.html
	前端地址:http://127.0.0.1:8080
`, global.Version, address)

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
