package system

import (
	"fmt"
	"net/url"
	"os"

	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pkg/errors"
)

var AutoCode = new(autoCode)

type autoCode struct{}

// GetDbs
// @Tags SystemAutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{} "获取成功!"
// @Router /autoCode/getDatabase [get]
func (a *autoCode) GetDbs(r *ghttp.Request) *response.Response {
	dbs, err := system.AutoCode.GetDbs()
	if err != nil {
		return &response.Response{Error: err, Message: "获取所有数据库失败!"}
	}
	return &response.Response{Data: g.Map{"dbs": dbs}, Message: "获取成功!"}
}

// GetTables
// @Tags SystemAutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param dbName query string true "数据库名"
// @Success 200 {object} response.Response{data=[]response.Table} "获取成功!"
// @Router /autoCode/getTables [get]
func (a *autoCode) GetTables(r *ghttp.Request) *response.Response {
	dbName := r.GetString("dbName", global.Config.Gorm.Dsn.GetDefaultDbName())
	tables, err := system.AutoCode.GetTables(dbName)
	if err != nil {
		return &response.Response{Error: err, Message: "查询table失败!"}
	}
	return &response.Response{Data: g.Map{"tables": tables}, Message: "获取成功!"}
}

// GetColumns
// @Tags SystemAutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param dbName query string true "数据库名"
// @Param tableName query string true "表名"
// @Success 200 {object} response.Response{data=[]response.Column} "获取成功!"
// @Router /autoCode/getColumn [get]
func (a *autoCode) GetColumns(r *ghttp.Request) *response.Response {
	dbName := r.GetString("dbName", global.Config.Gorm.Dsn.GetDefaultDbName())
	tableName := r.GetString("tableName")
	columns, err := system.AutoCode.GetColumns(tableName, dbName)
	if err != nil {
		return &response.Response{Error: err, Message: "获取表所有字段失败!"}
	}
	return &response.Response{Data: g.Map{"columns": columns}, Message: "获取成功!"}
}

// Preview
// @Tags SystemAutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoCodeCreate true "请求参数"
// @Success 200 {object} response.Response{} "获取成功!"
// @Router /autoCode/preview [post]
func (a *autoCode) Preview(r *ghttp.Request) *response.Response {
	var info request.AutoCodeCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "预览创建代码失败!"}
	}
	data, err := system.AutoCode.Preview(&info)
	if err != nil {
		return &response.Response{Error: err, Message: "预览创建代码失败!"}
	}
	return &response.Response{Data: g.Map{"autoCode": data}, Message: "预览成功!"}
}

// Create
// @Tags SystemAutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoCodeCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /autoCode/createTemp [post]
func (a *autoCode) Create(r *ghttp.Request) *response.Response {
	var info request.AutoCodeCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	err := system.AutoCode.Create(&info)
	if err != nil {
		_err := errors.New("创建代码成功并移动文件成功!")
		if errors.As(err, &_err) {
			r.Response.Header().Add("success", "true")
			r.Response.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./gf-vue-admin.zip")
		} else {
			r.Response.Header().Add("success", "false")
			r.Response.Header().Add("msg", url.QueryEscape(err.Error()))
		}
	} else {
		r.Response.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "gf-vue-admin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		r.Response.Header().Add("Content-Type", "application/json")
		r.Response.Header().Add("success", "true")
		r.Response.ServeFile("./gf-vue-admin.zip")
		_ = os.Remove("./gf-vue-admin.zip")
	}
	return &response.Response{}
}
