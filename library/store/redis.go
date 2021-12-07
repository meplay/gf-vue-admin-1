package store

import (
	"context"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewRedisStore() *_redis {
	return &_redis{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type _redis struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (r *_redis) UseWithCtx(ctx context.Context) base64Captcha.Store {
	r.Context = ctx
	return r
}

func (r *_redis) Set(id string, value string) error {
	return global.Redis.Set(r.Context, r.PreKey+id, value, r.Expiration).Err()
}

func (r *_redis) Get(key string, clear bool) string {
	val, err := global.Redis.Get(r.Context, key).Result()
	if err != nil {
		zap.L().Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err = global.Redis.Del(r.Context, key).Err()
		if err != nil {
			zap.L().Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (r *_redis) Verify(id, answer string, clear bool) bool {
	key := r.PreKey + id
	v := r.Get(key, clear)
	return v == answer
}
