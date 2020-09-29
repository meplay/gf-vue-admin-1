package v1

import (
	"fmt"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

func EmailTest(r *ghttp.Request) {
	err := service.EmailTest()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("发送失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, "发送成功")
}
