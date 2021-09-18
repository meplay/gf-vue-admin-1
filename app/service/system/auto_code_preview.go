package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Preview 预览创建代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) Preview(autoCode *system.AutoCodeStruct) (map[string]string, error) {
	dataList, _, needMkdir, err := s.getNeedList(autoCode)
	if err != nil {
		return nil, err
	}


	if err = utils.Directory.Creates(needMkdir...); err != nil {
		return nil, err
	} // 写入文件前，先创建文件夹

	// 创建map
	ret := make(map[string]string)

	// 生成map
	for _, value := range dataList {
		ext := ""
		if ext = filepath.Ext(value.AutoCodePath); ext == ".txt" {
			continue
		}
		f, openErr := os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755)
		if openErr != nil {
			return nil, openErr
		}
		if err = value.Template.Execute(f, autoCode); err != nil {
			return nil, err
		}
		_ = f.Close()
		f, err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		builder := strings.Builder{}
		builder.WriteString("```")

		if ext != "" && strings.Contains(ext, ".") {
			builder.WriteString(strings.Replace(ext, ".", "", -1))
		}
		builder.WriteString("\n\n")
		data, readErr := ioutil.ReadAll(f)
		if readErr != nil {
			return nil, readErr
		}
		builder.Write(data)
		builder.WriteString("\n\n```")

		pathArr := strings.Split(value.AutoCodePath, string(os.PathSeparator))
		ret[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = f.Close()

	}
	defer func() {
		if err = os.RemoveAll(autoPath); err != nil {
			zap.L().Error("移除中间文件失败!", zap.Error(err))
		}
	}()
	return ret, nil
}
