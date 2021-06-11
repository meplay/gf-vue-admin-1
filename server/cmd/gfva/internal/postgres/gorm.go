package postgres

import "flipped-aurora/gf-vue-admin/server/library/global"

var Postgres = new(_postgres)

type _postgres struct{}

func (p *_postgres) CreateDatabase() error {
	return global.Db.Exec("CREATE DATABASE "+global.Config.Gorm.Dsn.Sources[0].DbName).Error
}
