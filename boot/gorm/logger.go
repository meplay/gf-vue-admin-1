package boot

import (
	"fmt"

	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	if global.Config.Gorm.LogZap {
		zap.L().Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
