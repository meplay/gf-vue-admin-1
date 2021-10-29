package email

import (
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Api = new(api)

type api struct{}

// Test
// @Tags PluginEmail
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (a *api) Test(r *ghttp.Request) *response.Response {
	if err := Service.Test(); err != nil {
		return &response.Response{Error: err, Message: "发送邮件失败!"}
	}
	return &response.Response{Message: "发送邮件成功!"}
}

// Send
// @Tags PluginEmail
// @Summary 发送邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body Request true "请求参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/sendEmail [post]
func (a *api) Send(r *ghttp.Request) *response.Response {
	var info Request
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "发送失败!"}
	}
	if err := Service.Send(&info); err != nil {
		return &response.Response{Error: err, Message: "发送失败!"}
	}
	return &response.Response{Message: "发送成功!"}
}
