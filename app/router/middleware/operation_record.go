package middleware

import (
	"bytes"
	"io"
	"strconv"
	"time"

	model "github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/zap"
)

func OperationRecord(r *ghttp.Request) {
	// Request
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		zap.L().Error("读取内容失败", zap.Error(err))
	}

	r.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	id, _ := strconv.Atoi(r.Request.Header.Get("x-user-id"))

	record := request.OperationRecordCreate{
		OperationRecord: model.OperationRecord{
			Ip:      r.GetClientIp(),
			Method:  r.Request.Method,
			Path:    r.Request.URL.Path,
			Agent:   r.Request.UserAgent(),
			Request: string(body),
			UserID:  id,
		},
	}
	now := time.Now()

	r.Middleware.Next()

	// Response

	latency := time.Now().Sub(now)

	if err = r.GetError(); err != nil {
		record.ErrorMessage = err.Error()
	}

	record.Status = r.Response.Status
	record.Latency = time.Duration(latency.Microseconds())
	record.Response = string(r.Response.Buffer())

	if err = system.OperationRecord.Create(&record); err != nil {
		zap.L().Error("创建日志记录失败!", zap.Error(err))
	}
}
