package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var AutoCode = new(autoCode)

type autoCode struct{}

const (
	autoPath = "autocode_template/"
	basePath = "resource/template"
)

var injectionPaths []injectionMeta

func Init() {
	if len(injectionPaths) != 0 {
		return
	}
	injectionPaths = []injectionMeta{
		{
			path:        filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Boot, "gorm.go"),
			funcName:    "Initialize",
			structNameF: "new(example.%s),",
		},
		{
			path:        filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Boot, "router.go"),
			funcName:    "Initialize",
			structNameF: "example.New%sRouter(private).Private()",
		},
	}
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // 带格式化的
}

// GetAllTemplateFile 获取 pathName 文件夹下所有 tpl 文件
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetAllTemplateFile(path string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	length := len(files)
	fileList = make([]string, 0, length)
	for i := 0; i < length; i++ {
		if files[i].IsDir() {
			fileList, err = s.GetAllTemplateFile(path+"/"+files[i].Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(files[i].Name(), ".tpl") {
				fileList = append(fileList, path+"/"+files[i].Name())
			}
		}
	}
	return fileList, err
}

func (s *autoCode) DropTable(tableName string) error {
	return global.Db.Migrator().DropTable(tableName)
}

func (s *autoCode) getNeedList(info *system.AutoCodeStruct) (dataList []response.Template, fileList []string, needMkdir []string, err error) {
	info.TrimSpace() // 去除所有空格

	templateFileList, err := s.GetAllTemplateFile(basePath, nil)
	if err != nil {
		return nil, nil, nil, err
	} // 获取 basePath 文件夹下所有tpl文件

	length := len(templateFileList)
	dataList = make([]response.Template, 0, length)
	fileList = make([]string, 0, length)
	needMkdir = make([]string, 0, length) // 当文件夹下存在多个tpl文件时，改为map更合理

	for i := 0; i < length; i++ {
		dataList = append(dataList, response.Template{LocationPath: templateFileList[i]}) // 根据文件路径生成 tplData 结构体，待填充数据

		dataList[i].Template, err = template.ParseFiles(dataList[i].LocationPath)         // 生成 *Template, 填充 template 字段
		if err != nil {
			return nil, nil, nil, err
		}

		// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
		// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
		// resource/template/readme.txt.tpl -> autoCode/readme.txt
		trimBase := strings.TrimPrefix(dataList[i].LocationPath, basePath+"/")
		if trimBase == "readme.txt.tpl" {
			dataList[i].AutoCodePath = autoPath + "readme.txt"
			continue
		}

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				var fileName string
				if origFileName[firstDot:] != ".go" {
					fileName = info.PackageName + origFileName[firstDot:]
				} else {
					fileName = info.HumpPackageName + origFileName[firstDot:]
				}

				dataList[i].AutoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator], info.PackageName,
					origFileName[:firstDot], fileName)
			}
		}

		if lastSeparator := strings.LastIndex(dataList[i].AutoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[i].AutoCodePath[:lastSeparator])
		}

		fileList = append(fileList, dataList[i].AutoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

// injectionCode 封装代码注入
func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		if err := utils.Injection.AutoCode(meta.path, meta.funcName, code); err != nil {
			return err
		}
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}
