package service

import (
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/library/config"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/utils"
)

var System = new(_system)

type _system struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 读取配置文件
func (s *_system) GetConfig() *config.Config {
	return &global.Config
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置配置文件
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
