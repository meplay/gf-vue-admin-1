package utils

import (
	"os"

	"github.com/gogf/gf/frame/g"
)

var Directory = new(_directory)

type _directory struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 文件目录是否存在
func (d *_directory) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量创建文件夹
func (d *_directory) BatchCreate(directories ...string) error {
	for _, directory := range directories {
		if exist, err := d.PathExists(directory); err != nil {
			return err
		} else {
			if !exist {
				if err = os.MkdirAll(directory, os.ModePerm); err != nil {
					g.Log().Info("Function os.MkdirAll Failed!", g.Map{"err": err})
				}
			}
		}
	}
	g.Log().Info("Batch Create Succeed!", g.Map{"directory": directories})
	return nil
}
