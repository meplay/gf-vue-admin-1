package utils

import (
	"go.uber.org/zap"
	"os"
)

var Directory = new(directory)

type directory struct{}

// PathExists 文件目录是否存在
// Author [SliverHorn](https://github.com/SliverHorn)
func (d *directory) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

// Creates 批量创建文件夹
// Author [SliverHorn](https://github.com/SliverHorn)
func (d *directory) Creates(dirs ...string) error {
	length := len(dirs)
	for i := 0; i < length; i++ {
		exist, err := d.PathExists(dirs[i])
		if err != nil {
			return err
		}
		if !exist {
			zap.L().Info("正在创建文件夹中......", zap.String("文件夹名称`", dirs[i]))
			if err = os.MkdirAll(dirs[i], os.ModePerm); err != nil {
				zap.L().Error("创建文件夹失败!", zap.String("文件夹名称`", dirs[i]), zap.Error(err))
				return err
			}
			zap.L().Info("创建文件夹成功!", zap.String("文件夹名称`", dirs[i]))
		}
	}
	return nil
}
