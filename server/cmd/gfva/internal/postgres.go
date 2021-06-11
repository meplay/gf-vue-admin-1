//+build postgres

package internal

import (
	"fmt"
	"gf-vue-admin/boot"
	"gf-vue-admin/cmd/gfva/internal/postgres"
	"gf-vue-admin/cmd/gfva/internal/question"
	"gf-vue-admin/library/global"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"os"
)

var DbResolver = new(_postgres)

type _postgres struct {
	old string
	dsn string
	err error
}

func (p *_postgres) Error(s string, err error) {
	color.Warn.Printf("[postgres] --> %v! error:%v\n", s, err)
	os.Exit(0)
}

func (p *_postgres) SuccessFormat(format string, a ...interface{}) {
	color.LightGreen.Printf("[postgres] --> "+format, a)
}

func (p *_postgres) Database() {
	boot.DbResolver.Initialize()
	p.err = boot.DbResolver.Error()
	s := fmt.Sprintf("failed to connect to `host=127.0.0.1 user=root database=%v`: server error (FATAL: database \"%v\" does not exist (SQLSTATE 3D000))", global.Config.Gorm.Dsn.Sources[0].DbName, global.Config.Gorm.Dsn.Sources[0].DbName)
	if p.err.Error() == s {
		input := question.DatabaseAnswer{}
		if err := survey.Ask(question.DatabaseQuestion, &input); err != nil {
			p.Error("获取用户输入失败!", err)
		}
		switch input.Database {
		case "Link Start! gfva 为您创建数据库":
			p.old = global.Config.Gorm.Dsn.Sources[0].DbName
			global.Config.Gorm.Dsn.Sources[0].DbName = "postgres"
			boot.DbResolver.Initialize()
			if p.err = boot.DbResolver.Error(); p.err != nil {
				p.Error("链接数据库失败!", p.err)
			}
			global.Config.Gorm.Dsn.Sources[0].DbName = p.old
			if err := postgres.Postgres.CreateDatabase(); err != nil {
				p.Error("gfva 为您创建数据库失败!", err)
			}
			p.SuccessFormat("创建 %s 数据库成功!\n", global.Config.Gorm.Dsn.Sources[0].DbName)
		case "闪开!我自己来":
			os.Exit(0)
		case "退出程序":
			os.Exit(0)
		default:
			os.Exit(0)
		}
	}
}
