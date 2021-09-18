package response

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Column struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}

type Template struct {
	Template         *template.Template
	LocationPath     string
	AutoCodePath     string
	AutoMoveFilePath string
}

func (r *Template) GenerateAutoMoveFilePath() {
	base := filepath.Base(r.AutoCodePath)
	fileSlice := strings.Split(r.AutoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root,
				global.Config.AutoCode.Server.Router, base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Api, base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Service, base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Model, base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Request, base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root,
				global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Api, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Form, filepath.Base(filepath.Dir(filepath.Dir(r.AutoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			r.AutoMoveFilePath = filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Web.Root, global.Config.AutoCode.Web.Table, filepath.Base(filepath.Dir(filepath.Dir(r.AutoCodePath))), base)
		}
	}
}
