package boot

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var Redis = new(_redis)

type _redis struct{}

func (r *_redis) Initialize() {
	__redis := redis.NewClient(&redis.Options{
		DB:       global.Config.Redis.DB, // use default DB
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password, // no password set
	})
	pong, err := __redis.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("redis connect ping failed!", zap.Error(err))
		return
	}
	zap.L().Error("redis connect ping response", zap.String("ping", pong))
	global.Redis = __redis
}
