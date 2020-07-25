package boot

import (
	"server/library/global"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func InitializeRedis() {
	if g.Cfg().GetBool("system.UseMultipoint") {
		global.GFVA_REDIS = g.Redis()
	}
	Ping()
}

func Ping() {
	conn, err := global.GFVA_REDIS.Do("PING")
	if err != nil {
		g.Log().Error(err)
	}
	g.Log().Infof("redis connect ping response:%v", gconv.String(conn))
}
