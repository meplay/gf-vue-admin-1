package service

import (
	"flipped-aurora/gf-vue-admin/server/app/model/system/response"
	"flipped-aurora/gf-vue-admin/server/library/config"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"flipped-aurora/gf-vue-admin/server/library/utils"
)

var System = new(_system)

type _system struct{}

// GetConfig 读取配置文件
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *_system) GetConfig() *config.Config {
	return &global.Config
}

// SetConfig 设置配置文件
// Author: [SliverHorn](https://github.com/SliverHorn)
func (s *_system) SetConfig(info *config.Config) (err error) {
	configMap := utils.StructToMap(info)
	for k, v := range configMap {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}

func (s *_system) GetServerInfo() (*response.Server, error) {
	return utils.Server.Data()
}
