package internal

import (
	"fmt"
	"time"
)

var Upload = new(upload)

type upload struct{}

func (u *upload) GetObjectName(filename string) string {
	folder := time.Now().Format("20060102")
	return fmt.Sprintf("%s/%d%s", folder, time.Now().Unix(), filename) // 文件名格式 自己可以改 建议保证唯一性
}
