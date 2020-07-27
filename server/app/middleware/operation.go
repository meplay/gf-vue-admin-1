package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"server/app/api/request"
	v1 "server/app/api/v1"
	"server/app/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func OperationRecord(r *ghttp.Request) {
	var (
		body  []byte
		token *jwt.Token
		err   error
	)
	if r.Request.Method != http.MethodGet {
		body, err = ioutil.ReadAll(r.Request.Body)
		if err != nil {
			g.Log().Error("read body from request error:", err)
		} else {
			r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
	token, err = v1.GfJWTMiddleware.ParseToken(r)
	claims := gconv.Map(token.Claims)
	record := request.CreateOperation{
		Ip:      r.GetClientIp(),
		Method:  r.Request.Method,
		Path:    r.Request.URL.Path,
		Agent:   r.Request.UserAgent(),
		Request: string(body),
		UserId:  gconv.Int(claims["user_id"]),
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
}
