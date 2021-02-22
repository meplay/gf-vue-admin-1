package upload

import (
	"gf-vue-admin/interfaces"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func Oss() interfaces.Oss {
	switch global.Config.System.OssType {
	case "local":
		return Local
	case "qiniu":
		return Qiniu
	case "minio":
		if err := Minio.init(); err != nil {
			g.Log().Error("function Minio.init() Failed!", g.Map{"err": err})
		}
		return Minio
	case "aliyun":
		if err := AliYun.init(); err != nil {
			g.Log().Error("function AliYun.init() Failed!", g.Map{"err": err})
		}
		return AliYun
	default:
		return Local
	}
}
