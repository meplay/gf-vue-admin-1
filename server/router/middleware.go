package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 允许接口跨域请求
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 处理panic产生的错误
func Error(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log("exception").Error(err) // 记录到自定义错误日志文件
		r.Response.ClearBuffer()      //返回固定的友好信息
		//global.FailWithMessage(r, "服务器居然开小差了，请稍后再试吧！")
	}
}
