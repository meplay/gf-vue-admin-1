package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var JwtBlacklist = new(jwtBlacklist)

type jwtBlacklist struct{}

// JsonInBlacklist
// @Tags SystemJwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{} "拉黑成功!"
// @Router /jwt/jsonInBlacklist [post]
func (a *jwtBlacklist) JsonInBlacklist(r *ghttp.Request) *response.Response {
	token := r.Request.Header.Get("x-token")
	if err := system.JwtBlacklist.JwtToBlacklist(token); err != nil {
		return &response.Response{Error: err, Message: "拉黑失败!"}
	}
	return &response.Response{Message: "拉黑成功!"}
}
