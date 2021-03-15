package api

import (
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	service "gf-vue-admin/app/service/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Base = new(base)

type base struct{}

// @Tags SystemBase
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (b *base) Captcha(r *ghttp.Request) *response.Response {
	if result, err := service.Base.Captcha(); err != nil {
		return &response.Response{Data: result, MessageCode: response.ErrorCaptcha}
	} else {
		return &response.Response{Data: result, MessageCode: response.SuccessCaptcha}
	}
}

// @Tags SystemInitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
func (b *base) Initdb(r *ghttp.Request) *response.Response {
	if global.Db != nil {
		return &response.Response{Code: 7, Message: "非法访问!"}
	}
	var info request.InitDB
	if err := r.Parse(&info); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "参数校验不通过!"}
	}
	if err := service.Base.InitDB(&info); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "自动创建数据库失败，请查看后台日志!"}
	}
	return &response.Response{Code: 0, Message: "自动创建数据库成功!"}
}

// @Tags SystemCheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
func (b *base) Checkdb(r *ghttp.Request) *response.Response {
	if global.Db != nil {
		return &response.Response{Code: 0, Data: g.Map{"needInit": false}, Message: "数据库无需初始化!"}
	}
	return &response.Response{Code: 0, Data: g.Map{"needInit": true}, Message: "前往初始化数据库!"}
}
