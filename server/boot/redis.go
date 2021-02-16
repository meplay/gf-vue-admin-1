package boot

import (
	"context"
	"gf-vue-admin/library/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var Redis = new(_redis)

type _redis struct{}

func (r *_redis) Initialize() {
	client := redis.NewClient(&redis.Options{
		DB:       global.Config.Redis.DB, // use default DB
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password, // no password set
	})
	if pong, err := client.Ping(context.Background()).Result(); err != nil {
		zap.L().Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		zap.L().Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
}
