package response

import (
	"github.com/gogf/gf/net/ghttp"
)

type Response struct {
	Code        int         `json:"code"`
	MessageCode Code        `json:"-"`
	Data        interface{} `json:"data"`
	Error       error       `json:"-"`
	Err         string      `json:"err,omitempty"`
	Message     string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

type empty struct{}
type Handler struct{}

type handler func(r *ghttp.Request) *Response

func (h *Handler) Handler() func(handler handler) func(r *ghttp.Request) {
	return func(handler handler) func(r *ghttp.Request) {
		return func(r *ghttp.Request) {
			response := handler(r)
			if response.Data == nil {
				response.Data = empty{}
			}
			switch {
			case SuccessStart < response.MessageCode && response.MessageCode < SuccessEnd:
				response.Code = SUCCESS
				response.Message = response.MessageCode.Message()
				if err := r.Response.WriteJson(response); err != nil {
					panic(err)
				}
			case ErrorStart < response.MessageCode && response.MessageCode < ErrorEnd:
				response.Code = ERROR
				response.Message = response.MessageCode.Message()
				if response.Error != nil {
					response.Err = response.Error.Error()
				}
				if err := r.Response.WriteJson(response); err != nil {
					panic(err)
				}
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
