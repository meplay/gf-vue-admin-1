//+build postgres

package boot

import (
	extra "flipped-aurora/gf-vue-admin/server/app/model/extra"
	system "flipped-aurora/gf-vue-admin/server/app/model/system"
	workflow "flipped-aurora/gf-vue-admin/server/app/model/workflow"
	"flipped-aurora/gf-vue-admin/server/boot/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
)

var DbResolver = new(_postgres)

type _postgres struct {
	db  *gorm.DB
	dsn string
	err error
}

func (p *_postgres) Initialize() {
	resolver := p.GetResolver()
	p.db, p.err = gorm.Open(postgres.Open(p.dsn), internal.Gorm.GenerateConfig())
	if p.err != nil {
		g.Log().Error("postgres 链接失败!", g.Map{"err": p.err})
		return
	}
	defer func() {
		global.Db = p.db
	}()
	p.err = p.db.Use(resolver)
	if p.err != nil {
		g.Log().Error("mysql 链接集群失败!", g.Map{"err": p.err})
		return
	}
	if global.Config.Gorm.AutoMigrate {
		p.AutoMigrate()
	}
	sql, _ := p.db.DB()
	sql.SetMaxIdleConns(global.Config.Gorm.GetMaxIdleConnes())
	sql.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
}

func (p *_postgres) GetSources() []gorm.Dialector {
	length := len(global.Config.Gorm.Dsn.Sources)
	directories := make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		dsn := global.Config.Gorm.Dsn.Sources[i].GetPostgresDsn(global.Config.Gorm.Config)
		if i == 0 {
			p.dsn = dsn
		}
		directories = append(directories, postgres.Open(dsn))
	}
	return directories
}

func (p *_postgres) GetReplicas() []gorm.Dialector {
	length := len(global.Config.Gorm.Dsn.Replicas)
	directories := make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		dsn := global.Config.Gorm.Dsn.Replicas[i].GetPostgresDsn(global.Config.Gorm.Config)
		directories = append(directories, postgres.Open(dsn))
	}
	return directories
}

func (p *_postgres) GetResolver() gorm.Plugin {
	sources := p.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: p.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.Config.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.Config.Gorm.GetConnMaxLifetime())
	return resolver
}

func (p *_postgres) AutoMigrate() {
	err := p.db.AutoMigrate(
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
	g.Log().Info(`注册表成功!`)
}

func (p *_postgres) Error() error {
	return p.err
}
