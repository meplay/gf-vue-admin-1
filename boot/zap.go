package boot

import (
	"fmt"

	"github.com/flipped-aurora/gf-vue-admin/boot/internal"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"go.uber.org/zap"
)

var Zap = new(_zap)

type _zap struct {
	zap *zap.Logger
}

func (z *_zap) Initialize() {
	// todo 判断是否有Director文件夹

	level := global.Config.Zap.GetZapLevel()

	writer, err := internal.Zap.GetWriteSyncer()
	if err != nil {
		fmt.Println(`获取WriteSyncer失败, err: `, err)
		return
	} // 使用 file-rotatelogs 进行日志分割

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		z.zap = zap.New(internal.Zap.GetEncoderCore(writer, level), zap.AddStacktrace(level))
	} else {
		z.zap = zap.New(internal.Zap.GetEncoderCore(writer, level))
	}

	if global.Config.Zap.ShowLine {
		z.zap = z.zap.WithOptions(zap.AddCaller())
	}

	zap.ReplaceGlobals(z.zap)
}
