package service

import (
	"server/app/model"
)

// GetSystemConfig Get configuration information
// GetSystemConfig 获取配置信息
func GetSystemConfig() (config *model.ServerConfig) {
	return &model.Config
}
