package v1

import (
	"fmt"
	"net/url"
	"os"
	"server/app/api/request"
	"server/app/model"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CreateTemp Automatic code templates
// CreateTemp 自动代码模板
func CreateTemp(r *ghttp.Request) {
	var a model.AutoCodeStruct
	if err := r.Parse(&a); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if a.AutoCreateApiToSql {
		creates := [6]*request.CreateApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
				Description: "批量删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "PUT",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.Description + "列表",
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
		}
		for _, create := range creates {
			if err := service.CreateApi(create); err != nil {
				r.Header.Set("success", "false")
				r.Header.Set("msg", url.QueryEscape(fmt.Sprintf("自动化创建失败，%v，请自行清空垃圾数据", err)))
				r.Exit()
			}
		}
	}
	if err := service.CreateTemp(a); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		if err = os.Remove("./gf-vue-admin.zip"); err != nil {
			panic(err)
		}
	}
	r.Response.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "gf-vue-admin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	r.Response.Header().Add("Content-Type", "application/json")
	r.Response.Header().Add("success", "true")
	r.Response.ServeFile("./gf-vue-admin.zip")
	//r.Response.ServeFileDownload("./gf-vue-admin.zip", "gf-vue-admin.zip")
	if err := os.Remove("./gf-vue-admin.zip"); err != nil {
		panic(err)
	}
}

// GetTables Get the table of the corresponding database
// GetTables 获取对应数据库的表
func GetTables(r *ghttp.Request) {
	dbName := r.GetString("dbName", "")
	tables, err := service.GetTables(dbName)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询table失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"tables": tables})
}

// GetDB Get the database
// GetDB 获取数据库
func GetDB(r *ghttp.Request) {
	dbs, err := service.GetDB()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询table失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"dbs": dbs})
}

// GetColumns Get all field information of the specified table
// GetColumns 获取指定表所有字段信息
func GetColumns(r *ghttp.Request) {
	dbName := r.GetString("dbName", "")
	tableName := r.GetString("tableName")
	columns, err := service.GetColumn(dbName, tableName)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询table失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"columes": columns})
}
