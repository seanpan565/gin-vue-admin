// viper.go 加载并监听 config.yaml，支持命令行、环境变量与 gin 模式切换。
package core

import (
	"flag"
	"fmt"
	"os"

	"mall-admin/server/core/internal"
	"mall-admin/server/global"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper 读取配置文件并绑定到 global.GVA_CONFIG，支持热更新。
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if global.GVA_LOG != nil {
			global.GVA_LOG.Info("配置文件已变更", zap.String("file", e.Name))
		}
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			if global.GVA_LOG != nil {
				global.GVA_LOG.Error("热加载配置失败", zap.Error(err))
			}
			return
		}
		ApplyEnvOverrides()
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
	ApplyEnvOverrides()

	return v
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {
	// `-c` flag parse
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}
	if env := os.Getenv(internal.ConfigEnv); env != "" { // 判断环境变量 GVA_CONFIG
		config = env
		fmt.Printf("您正在使用 %s 环境变量, config 的路径为 %s\n", internal.ConfigEnv, config)
		return
	}

	switch gin.Mode() { // 根据 gin 模式文件名
	case gin.DebugMode:
		config = internal.ConfigDebugFile
	case gin.ReleaseMode:
		config = internal.ConfigReleaseFile
	case gin.TestMode:
		config = internal.ConfigTestFile
	}
	fmt.Printf("您正在使用 gin 的 %s 模式运行, config 的路径为 %s\n", gin.Mode(), config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
