package middleware

import (
	"bytes"
	"io/ioutil"
	"server/app/api/request"
	v1 "server/app/api/v1"
	"server/app/service"
	"server/library/utils"
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func ErrorToEmail(r *ghttp.Request) {
	var adminId int
	var username string
	body, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		g.Log().Error("read body from request error:", err)
	} else {
		r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	if token, err := v1.GfJWTMiddleware.ParseToken(r); err != nil { // 优先从jwt获取用户信息
		id, _ := strconv.Atoi(r.Request.Header.Get("x-user-id"))
		adminReturn, findErr := service.FindAdminById(id)
		if findErr != nil {
			g.Log().Errorf("err:%v", findErr.Error())
		}
		username = adminReturn.Username
		adminId = int(adminReturn.Id)
	} else {
		claims := gconv.Map(token.Claims)
		adminReturn, findError := service.FindAdmin(gconv.String(claims["admin_uuid"]))
		if findError != nil {
			g.Log().Errorf("err:%v", findError.Error())
		}
		username = adminReturn.Username
		adminId = int(adminReturn.Id)
	}

	record := request.CreateOperation{
		Ip:      r.GetClientIp(),
		Method:  r.Request.Method,
		Path:    r.Request.URL.Path,
		Agent:   r.Request.UserAgent(),
		Request: string(body),
		UserId:  adminId,
	}
	now := time.Now()
	// Request
	r.Middleware.Next()
	// Response
	latency := time.Now().Sub(now)
	if err := r.GetError(); err != nil {
		record.ErrorMessage = err.Error()
	}
	record.Status = r.Response.Status
	record.Latency = latency.Microseconds()
	record.Response = string(r.Response.Buffer())

	if err := service.CreateOperation(&record); err != nil {
		g.Log().Error("create operation record error:", err)
	}
	str := "接收到的请求为" + record.Request + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
	if g.Cfg("system").GetBool("system.ErrorToEmail") {
		if record.Status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				g.Log().Errorf("ErrorToEmail Failed, err:%v", err)
			}
		}
	}
}
