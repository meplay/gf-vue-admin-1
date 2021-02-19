package api

import (
	"gf-vue-admin/app/api/response"
	service "gf-vue-admin/app/service/system"
	"gf-vue-admin/library/config"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Config = new(_config)

type _config struct{}

// @Tags SystemConfig
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/getSystemConfig [post]
func (*_config) GetConfig(r *ghttp.Request) *response.Response {
	return &response.Response{Data: g.Map{"config": service.Config.GetConfig()}, MessageCode: response.SuccessOperation}
}

// @Tags SystemConfig
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body config.Server true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /system/setSystemConfig [post]
func (*_config) SetConfig(r *ghttp.Request) *response.Response {
	var info config.Config
	if err := service.Config.SetConfig(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorOperation}
	}
	return &response.Response{MessageCode: response.SuccessOperation}
}
