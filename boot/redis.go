package boot

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var Redis = new(__redis)

type __redis struct{}

func (r *__redis) Initialize() {
	_redis := redis.NewClient(&redis.Options{
		DB:       global.Config.Redis.DB, // use default DB
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password, // no password set
	})
	pong, err := _redis.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("redis connect ping failed!", zap.Error(err))
		return
	}
	zap.L().Error("redis connect ping response", zap.String("ping", pong))
	global.Redis = _redis
}
