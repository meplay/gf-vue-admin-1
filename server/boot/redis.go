package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func InitializeRedis() {
	if g.Cfg().GetBool("system.UseMultipoint") {
		conn, err := g.Redis().Do("PING")
		if err != nil {
			g.Log().Error(err)
		}
		g.Log().Infof("redis connect ping response:%v", gconv.String(conn))
	}
}
