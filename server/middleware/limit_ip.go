// limit_ip.go 基于 Redis 的 IP 访问频率限流中间件。
package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"mall-admin/server/global"
	"mall-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey func(c *gin.Context) string
	// 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey 默认生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	expiration := time.Duration(expire) * time.Second
	if global.GVA_REDIS != nil {
		if err = SetLimitWithTime(key, limit, expiration); err != nil {
			global.GVA_LOG.Error("limit", zap.Error(err))
		}
		return err
	}
	return memoryLimitWithTime(key, limit, expiration)
}

// memoryLimitWithTime Redis 不可用时使用进程内缓存限流。
func memoryLimitWithTime(key string, limit int, expiration time.Duration) error {
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, expiration)
		return nil
	}
	count, _ := v.(int)
	if count >= limit {
		return errors.New("请求太过频繁，请稍后再试")
	}
	global.BlackCache.Increment(key, 1)
	return nil
}

// DefaultLimit 按客户端 IP 限制单位时间内的请求次数。
func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.GVA_CONFIG.System.LimitTimeIP,
		Limit:         global.GVA_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}

// AuthRateLimit 登录/注册等认证接口的更严格限流。
func AuthRateLimit() gin.HandlerFunc {
	limit := global.GVA_CONFIG.Security.AuthLimitCount
	expire := global.GVA_CONFIG.Security.AuthLimitTime
	if limit <= 0 {
		limit = 30
	}
	if expire <= 0 {
		expire = 60
	}
	return LimitConfig{
		GenerationKey: func(c *gin.Context) string {
			return "GVA_AuthLimit:" + c.ClientIP() + ":" + c.FullPath()
		},
		CheckOrMark: DefaultCheckOrMark,
		Expire:      expire,
		Limit:       limit,
	}.LimitWithTime()
}

// SetLimitWithTime 设置访问次数
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.GVA_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// 次数
		if times, err := global.GVA_REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.GVA_REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.GVA_REDIS.Incr(context.Background(), key).Err()
			}
		}
	}
}
