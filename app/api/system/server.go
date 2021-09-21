package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Server = new(server)

type server struct{}

// ReloadSystem
// @Tags SystemServer
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{} "重启系统成功!"
// @Router /system/reloadSystem [post]
func (s *server) ReloadSystem(r *ghttp.Request) *response.Response {
	if err := utils.Server.Reload(); err != nil {
		return &response.Response{Error: err, Message: "重启系统失败!"}
	}
	return &response.Response{Message: "重启系统成功!"}
}

// GetServerInfo
// @Tags SystemServer
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]response.Server} "获取成功!"
// @Router /system/getServerInfo [post]
func (s *server) GetServerInfo(r *ghttp.Request) *response.Response {
	data, err := system.Server.GetServerInfo()
	if err != nil {
		return &response.Response{Error: err, Message: "获取失败!"}
	}
	return &response.Response{Data: g.Map{"server": data}, Message: "获取成功!"}
}

// GetSystemConfig
// @Tags SystemServer
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Success 200 {object} response.Response{data=config.Config} "获取成功!"
// @Router /system/getSystemConfig [post]
func (s *server) GetSystemConfig(r *ghttp.Request) *response.Response {
	return &response.Response{}
}

// SetSystemConfig
// @Tags SystemServer
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body config.Config true "设置配置文件内容"
// @Success 200 {object} response.Response{} "设置成功!"
// @Router /system/setSystemConfig [post]
func (s *server) SetSystemConfig(r *ghttp.Request) *response.Response {
	return &response.Response{}
}
