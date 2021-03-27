package boot

import (
	"database/sql"
	extra "gf-vue-admin/app/model/extra"
	system "gf-vue-admin/app/model/system"
	workflow "gf-vue-admin/app/model/workflow"
	"gf-vue-admin/boot/internal"
	"gf-vue-admin/library/gdbadapter"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DatabaseInfo struct {
	Value        string `gorm:"column:Value"`
	VariableName string `gorm:"column:Variable_name"`
}

var Mysql = &_mysql{_config: &gorm.Config{}}

type _mysql struct {
	db      *gorm.DB
	err     error
	sql     *sql.DB
	_config *gorm.Config

	old       string // 配置文件第一次读取数据库数据
	input     string
	version   string
	character string
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm连接mysql数据库
func (m *_mysql) Initialize() {
	if m.db, m.err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         191,                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                     // 根据版本自动配置
	}), internal.GenerateConfig()); m.err != nil {
		g.Log().Error(`Gorm连接MySQL异常!`, g.Map{"err": m.err})
	} else {
		if m.sql, m.err = m.db.DB(); m.err != nil {
			g.Log().Error(`DatabaseSql对象获取异常!`, g.Map{"err": m.err})
		} else {
			global.Db = m.db
			m.AutoMigrateTables()
			m.sql.SetMaxIdleConns(global.Config.Mysql.GetMaxIdleConnes())
			m.sql.SetMaxOpenConns(global.Config.Mysql.GetMaxOpenConnes())
		}
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm 同步模型 生成mysql表
func (m *_mysql) AutoMigrateTables() {
	if !global.Db.Migrator().HasTable("casbin_rule") {
		m.err = global.Db.Migrator().CreateTable(&gdbadapter.CasbinRule{})
	}
	m.err = global.Db.AutoMigrate(
		new(system.Api),
		new(system.Admin),
		new(system.Menu),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),

		new(extra.File),
		new(extra.SimpleUploader),
		new(extra.BreakpointContinue),
		new(extra.BreakpointContinueChunk),

		new(workflow.WorkflowNode),
		new(workflow.WorkflowMove),
		new(workflow.WorkflowEdge),
		new(workflow.WorkflowProcess),
		new(workflow.WorkflowEndPoint),
		new(workflow.WorkflowStartPoint),
	)
	if m.err != nil {
		g.Log().Error(`注册表失败!`, g.Map{"err": m.err})
		os.Exit(0)
	}
	g.Log().Info(`注册表成功!`)
}
