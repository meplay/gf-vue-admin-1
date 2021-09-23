package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Preview 预览创建代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) Preview(info *request.AutoCodeCreate) (map[string]string, error) {
	dataList, _, needMkdir, err := s.getNeedList(&info.AutoCodeStruct)
	if err != nil {
		return nil, err
	}

	if err = utils.Directory.Creates(needMkdir...); err != nil {
		return nil, err
	} // 写入文件前，先创建文件夹
	
	length := len(dataList)
	ret := make(map[string]string, length) // 创建map

	for i := 0; i < length; i++ {
		var ext string
		if ext = filepath.Ext(dataList[i].AutoCodePath); ext == ".txt" {
			continue
		}
		var file *os.File
		file, err = os.OpenFile(dataList[i].AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "打开文件失败!")
		}
		if err = dataList[i].Template.Execute(file, info.AutoCodeStruct); err != nil {
			return nil, err
		}
		_ = file.Close()
		file, err = os.OpenFile(dataList[i].AutoCodePath, os.O_CREATE|os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		builder := strings.Builder{}
		builder.WriteString("```")

		if ext != "" && strings.Contains(ext, ".") {
			builder.WriteString(strings.Replace(ext, ".", "", -1))
		}
		builder.WriteString("\n\n")
		data, readErr := ioutil.ReadAll(file)
		if readErr != nil {
			return nil, readErr
		}
		builder.Write(data)
		builder.WriteString("\n\n```")

		pathArr := strings.Split(dataList[i].AutoCodePath, string(os.PathSeparator))
		ret[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = file.Close()
	} // 生成map数据

	defer func() {
		if err = os.RemoveAll(autoPath); err != nil {
			zap.L().Error("移除中间文件失败!", zap.Error(err))
		}
	}()
	return ret, nil
}
