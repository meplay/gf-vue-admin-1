package response

import (
	"fmt"

	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/zap"
)

type Response struct {
	Code        int         `json:"code" swaggertype:"string" example:"int 状态码(成功:0, 失败:7)"`
	MessageCode Code        `json:"-"`
	Data        interface{} `json:"data" swaggertype:"string" example:"object 数据"`
	Error       error       `json:"-"`
	Err         string      `json:"err,omitempty" example:"错误信息"`
	Message     string      `json:"msg" example:"消息"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

type (
	empty   struct{}
	Handler struct{}
)

type handler func(r *ghttp.Request) *Response

func (h *Handler) Handler() func(handler handler) func(r *ghttp.Request) {
	return func(handler handler) func(r *ghttp.Request) {
		return func(r *ghttp.Request) {
			response := handler(r)
			if response.Data == nil {
				response.Data = empty{}
			}
			if response.Error != nil {
				response.Code = ERROR
				response.Err = response.Error.Error()
				zap.L().Error(fmt.Sprintf("%+v", response.Error))
			}
			switch {
			case SuccessStart < response.MessageCode && response.MessageCode < SuccessEnd:
				response.Code = SUCCESS
				response.Message = response.MessageCode.Message()
				_ = r.Response.WriteJson(response)
			case ErrorStart < response.MessageCode && response.MessageCode < ErrorEnd:
				response.Code = ERROR
				response.Message = response.MessageCode.Message()
				if response.Error != nil {
					response.Err = response.Error.Error()
				}
				_ = r.Response.WriteJson(response)
			default:
				if response.Error != nil {
					response.Err = response.Error.Error()
					response.Code = 7
					_ = r.Response.WriteJson(response)
				} else {
					response.Code = 0
					if response.Message == "" {
						response.Message = SuccessOperation.Message()
					}
					_ = r.Response.WriteJson(response)
				}
			}
		}
	}
}
