package service

import (
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/library/config"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/utils"
)

var Config = new(_config)

type _config struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 读取配置文件
func (c *_config) GetConfig() *config.Config {
	return &global.Config
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置配置文件
func (c *_config) SetConfig(info *config.Config) (err error) {
	configMap := utils.StructToMap(info)
	for k, v := range configMap {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}

func (c *_config) GetServerInfo() (*response.Server, error) {
	return utils.Server.Data()
}
