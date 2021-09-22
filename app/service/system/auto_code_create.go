package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

// CreateTemp 创建代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) CreateTemp(info *system.AutoCodeStruct) error {

	if info.AutoMoveFile && AutoCodeHistory.Repeat(info.StructName) {
		return errors.New("重复创建!")
	} // 增加判断: 重复创建struct

	dataList, fileList, needMkdir, err := s.getNeedList(info)
	if err != nil {
		return errors.Wrap(err, "获取模板信息失败!")
	}

	if err = utils.Directory.Creates(needMkdir...); err != nil {
		return errors.Wrap(err, "创建文件夹失败!")
	} // 写入文件前，先创建文件夹

	for _, value := range dataList {
		var file *os.File
		file, err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return errors.Wrap(err, "打开模板文件失败!")
		}
		if err = value.Template.Execute(file, info); err != nil {
			return errors.Wrap(err, "模板数据填充失败!")
		}
		_ = file.Close()
	} // 生成文件

	defer func() { // 移除中间文件
		if err = os.RemoveAll(autoPath); err != nil {
			zap.L().Error("移除中间文件失败!", zap.Error(err))
		}
	}()

	var apis system.AutoCodeApis
	apis, err = s.AutoCreateApi(info)
	if err != nil {
		return errors.Wrap(err, "创建api记录失败!")
	}

	if info.AutoMoveFile { // 判断是否需要自动转移
		length := len(dataList)
		for i := 0; i < length; i++ {
			dataList[i].GenerateAutoMoveFilePath() // 生成每个文件移动后的文件路径
			if err = utils.File.Move(dataList[i].AutoCodePath, dataList[i].AutoMoveFilePath); err != nil {
				return err
			} // 移动文件
		}

		entities := info.GenerateInjection() // 生成注入内容
		for i := 0; i < len(entities); i++ {
			if err = utils.Injection.AutoCode(entities[i].Filepath, entities[i].FunctionName, entities[i].CodeData); err != nil {
				return errors.Wrap(err, "注入代码失败!")
			}
		}

		if global.Config.AutoCode.Restart {
			go func() {
				_ = utils.Server.Reload()
			}()
		}
	} else { // 打包
		if err = utils.File.ZipFiles("./gf-vue-admin.zip", fileList, ".", "."); err != nil {
			return errors.Wrap(err, "打包失败!")
		}
	}

	if info.AutoMoveFile || info.AutoCreateApiToSql {
		if err = AutoCodeHistory.Create(dataList.ToRequestAutoCodeHistoryCreate(info, apis)); err != nil {
			return errors.Wrap(err, "创建自动化代码历史记录失败!")
		}
		return errors.New("创建代码成功并移动文件成功!")
	}
	return nil
}
