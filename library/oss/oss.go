package oss

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

func Oss() interfaces.Oss {
	switch global.Config.System.OssType {
	case "local":
		return Local
	case "qiniu":
		return Qiniu
	case "minio":
		return Minio
	case "aliyun":
		return Aliyun
	case "tencent":
		return Tencent
	default:
		return Local
	}
}
