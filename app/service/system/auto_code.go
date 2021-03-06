package system

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

var AutoCode = new(autoCode)

type autoCode struct{}

const (
	autoPath = "autocode_template/"
	basePath = "public/template"
)

type Templates []Template

type Template struct {
	Template         *template.Template
	LocationPath     string
	AutoCodePath     string
	AutoMoveFilePath string
}

func (t *Template) GenerateAutoMoveFilePath() {
	base := filepath.Base(t.AutoCodePath)
	fileSlice := strings.Split(t.AutoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root,
				global.Config.AutoCode.Server.Router, base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Api, base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Service, base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Model, base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Request, base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Api, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Form, filepath.Base(filepath.Dir(filepath.Dir(t.AutoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			t.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Table, filepath.Base(filepath.Dir(filepath.Dir(t.AutoCodePath))), base)
		}
	}
}

func (t *Templates) ToRequestAutoCodeHistoryCreate(info *system.AutoCodeStruct, apis system.AutoCodeApis) *request.AutoCodeHistoryCreate {
	if t == nil || info == nil {
		return nil
	}
	entity := request.AutoCodeHistoryCreate{
		Apis:           apis,
		Injection:      info.Injection,
		AutoCodeStruct: *info,
	}
	templates := *t
	length := len(templates)
	paths := make(system.AutoCodePaths, 0, length)
	for i := 0; i < length; i++ {
		if templates[i].AutoMoveFilePath == "" {
			continue
		}
		paths = append(paths, system.AutoCodePath{Filepath: templates[i].AutoMoveFilePath})
	}
	entity.AutoCodePaths = paths
	return &entity
}

// GetAllTemplateFile ?????? pathName ?????????????????? tpl ??????
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetAllTemplateFile(path string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	length := len(files)
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

func (s *autoCode) getNeedList(info *system.AutoCodeStruct) (dataList Templates, fileList []string, needMkdir []string, err error) {
	info.TrimSpace() // ??????????????????

	var templateFileList []string
	templateFileList, err = s.GetAllTemplateFile(basePath, nil)
	if err != nil {
		return nil, nil, nil, err
	} // ?????? basePath ??????????????????tpl??????

	length := len(templateFileList)
	dataList = make(Templates, 0, length)
	fileList = make([]string, 0, length)
	needMkdir = make([]string, 0, length) // ???????????????????????????tpl??????????????????map?????????

	for i := 0; i < length; i++ {
		dataList = append(dataList, Template{LocationPath: templateFileList[i]}) // ???????????????????????? tplData ???????????????????????????

		dataList[i].Template, err = template.ParseFiles(dataList[i].LocationPath) // ?????? *Template, ?????? template ??????
		if err != nil {
			return nil, nil, nil, err
		}

		// ??????????????????????????? autoCodePath ?????????readme.txt.tpl????????????????????????????????????
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
