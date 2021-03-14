package interfaces

import (
	"github.com/gookit/color"
)

type InitDateFunc interface {
	Init() error
	TableName() string
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量初始化表数据
func InitDb(inits ...InitDateFunc) error {
	for _, init := range inits {
		if err := init.Init(); err != nil {
			color.Warn.Printf("\n[Mysql] --> %v 表初始数据失败, err: %v\n", init.TableName(), err)
			return err
		}
	}
	return nil
}
