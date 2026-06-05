// Package captcha 提供验证码 Redis 存储后端。
package captcha

// 验证码 Redis 存储实现，满足 base64Captcha 存储接口。
import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

// NewDefaultRedisStore 创建默认配置的 Redis 验证码存储
func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}

// RedisStore 基于 Redis 的验证码存储
type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

// UseWithCtx 绑定上下文
func (rs *RedisStore) UseWithCtx(ctx context.Context) *RedisStore {
	if ctx != nil {
		rs.Context = ctx
	}
	return rs
}

// Set 存储验证码
func (rs *RedisStore) Set(id string, value string) error {
	err := global.GVA_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

// Get 获取验证码，可选验证后清除
func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.GVA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.GVA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

// Verify 校验验证码答案
func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
