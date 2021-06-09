//+build mysql

package boot

import (
	extra "gf-vue-admin/app/model/extra"
	system "gf-vue-admin/app/model/system"
	workflow "gf-vue-admin/app/model/workflow"
	"gf-vue-admin/boot/internal"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
)

var DbResolver = new(_mysql)

type _mysql struct {
	dsn string
}

// Initialize gorm连接mysql数据库
// Author SliverHorn
func (m *_mysql) Initialize() {
	resolver := m.GetResolver()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.dsn, // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: true,  // 根据版本自动配置
	}), internal.Gorm.GenerateConfig())
	if err != nil {
		g.Log().Error(`mysql 链接失败!`, g.Map{"err": err})
		os.Exit(0)
	}
	err = db.Use(resolver)
	if err != nil {
		g.Log().Error("mysql 链接集群失败!", g.Map{"err": err})
		os.Exit(0)
	}
	global.Db = db
	if global.Config.Gorm.AutoMigrate {
		m.AutoMigrate()
	}
	sql, _ := db.DB()
	sql.SetMaxIdleConns(global.Config.Gorm.GetMaxIdleConnes())
	sql.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
}

func (m *_mysql) GetSources() []gorm.Dialector {
	length := len(global.Config.Gorm.Dsn.Sources)
	directories := make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		dsn := global.Config.Gorm.Dsn.Sources[i].GetMysqlDsn(global.Config.Gorm.Config)
		if i == 0 {
			m.dsn = dsn
		}
		directories = append(directories, mysql.Open(dsn))
	}
	return directories
}

func (m *_mysql) GetReplicas() []gorm.Dialector {
	length := len(global.Config.Gorm.Dsn.Replicas)
	directories := make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		dsn := global.Config.Gorm.Dsn.Replicas[i].GetMysqlDsn(global.Config.Gorm.Config)
		directories = append(directories, mysql.Open(dsn))
	}
	return directories
}

func (m *_mysql) GetResolver() gorm.Plugin {
	sources := m.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: m.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.Config.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.Config.Gorm.GetConnMaxLifetime())
	return resolver
}

func (m *_mysql) AutoMigrate() {
	err := global.Db.AutoMigrate(
		new(system.Api),
		new(system.Menu),
		new(system.Admin),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),

		new(workflow.WorkflowNode),
		new(workflow.WorkflowMove),
		new(workflow.WorkflowEdge),
		new(workflow.WorkflowLeave),
		new(workflow.WorkflowProcess),
		new(workflow.WorkflowEndPoint),
		new(workflow.WorkflowStartPoint),

		new(extra.File),
		new(extra.SimpleUploader),
		new(extra.BreakpointContinue),
		new(extra.BreakpointContinueChunk),
	)
	if err != nil {
		g.Log().Error(`注册表失败!`, g.Map{"err": err})
		os.Exit(0)
	}
	zap.L().Info(`注册表成功!`)
}
