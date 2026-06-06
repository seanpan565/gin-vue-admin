// redis.go 初始化 Redis 单例及多实例连接池。
package initialize

import (
	"context"

	"mall-admin/server/config"
	"mall-admin/server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(redisCfg config.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.String("name", redisCfg.Name), zap.Error(err))
		return nil, err
	}

	global.GVA_LOG.Info("redis connect ping response:", zap.String("name", redisCfg.Name), zap.String("pong", pong))
	return client, nil
}

// Redis 初始化默认 Redis 客户端并写入 global.GVA_REDIS。
func Redis() {
	redisClient, err := initRedisClient(global.GVA_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.GVA_REDIS = redisClient
}

// RedisList 初始化多 Redis 实例，供多点登录等场景使用。
func RedisList() {
	redisMap := make(map[string]redis.UniversalClient)

	for _, redisCfg := range global.GVA_CONFIG.RedisList {
		client, err := initRedisClient(redisCfg)
		if err != nil {
			panic(err)
		}
		redisMap[redisCfg.Name] = client
	}

	global.GVA_REDISList = redisMap
}
