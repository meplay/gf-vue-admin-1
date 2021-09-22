package system

import (
	"encoding/json"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

// CreateTemp 创建代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) CreateTemp(autoCode *system.AutoCodeStruct, ids ...uint) error {

	if autoCode.AutoMoveFile && AutoCodeHistory.Repeat(autoCode.StructName) {
		return errors.New("重复创建!")
	} // 增加判断: 重复创建struct

	dataList, fileList, needMkdir, err := s.getNeedList(autoCode)
	if err != nil {
		return err
	}

	meta, _ := json.Marshal(autoCode)

	if err = utils.Directory.Creates(needMkdir...); err != nil {
		return err
	} // 写入文件前，先创建文件夹

	for _, value := range dataList {
		file, _err := os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if _err != nil {
			return _err
		}
		if err = value.Template.Execute(file, autoCode); err != nil {
			return err
		}
		_ = file.Close()
	} // 生成文件

	defer func() { // 移除中间文件
		if err = os.RemoveAll(autoPath); err != nil {
			zap.L().Error("移除中间文件失败!", zap.Error(err))
		}
	}()
	bf := strings.Builder{}
	idBf := strings.Builder{}
	injectionCodeMeta := strings.Builder{}
	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}
	if autoCode.AutoMoveFile { // 判断是否需要自动转移
		Init()
		length := len(dataList)
		for i := 0; i < length; i++ {
			dataList[i].GenerateAutoMoveFilePath() // 生成每个文件移动后的文件路径
			if err = utils.File.Move(dataList[i].AutoCodePath, dataList[i].AutoMoveFilePath); err != nil {
				return err
			} // 移动文件
		}

		err = injectionCode(autoCode.StructName, &injectionCodeMeta)
		if err != nil {
			return err
		}
		// 保存生成信息
		for _, data := range dataList {
			if len(data.AutoMoveFilePath) != 0 {
				bf.WriteString(data.AutoMoveFilePath)
				bf.WriteString(";")
			}
		}

		if global.Config.AutoCode.Restart {
			go func() {
				_ = utils.Server.Reload()
			}()
		}
	} else { // 打包
		if err = utils.File.ZipFiles("./gf-vue-admin.zip", fileList, ".", "."); err != nil {
			return err
		}
	}
	if autoCode.AutoMoveFile || autoCode.AutoCreateApiToSql {
		if autoCode.TableName != "" {
			err = AutoCodeHistory.Create(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.TableName,
				idBf.String(),
			)
		} else {
			err = AutoCodeHistory.Create(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.StructName,
				idBf.String(),
			)
		}
	}
	if err != nil {
		return err
	}
	if autoCode.AutoMoveFile {
		return errors.New("创建代码成功并移动文件成功!")
	}
	return nil

}
