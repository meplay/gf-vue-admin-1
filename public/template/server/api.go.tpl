package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/example"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct{}

// Create
// @Tags Example{{.StructName}}
// @Summary 创建 {{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.{{.StructName}}Create true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func (s *{{.Abbreviation}}) Create(r *ghttp.Request) *response.Response {
	var info request.{{.StructName}}Create
    if err := r.Parse(&info); err != nil {
        return &response.Response{Error: err, MessageCode: response.ErrorCreated}
    }
    if err := example.{{.StructName}}.Create(&info); err != nil {
        return &response.Response{Error: err, MessageCode: response.ErrorCreated}
    }
    return &response.Response{MessageCode: response.SuccessCreated}
}

// Find
// @Tags Example{{.StructName}}
// @Summary 用id查询 {{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "获取数据成功!"
// @Router /{{.Abbreviation}}/find [get]
func (s *{{.Abbreviation}}) Find(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFind}
	}
	data, err := example.{{.StructName}}.Find(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFind}
	}
	return &response.Response{Data: g.Map{"{{.Abbreviation}}": data}, MessageCode: response.SuccessFind}
}

// Update
// @Tags Example{{.StructName}}
// @Summary 更新 {{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.{{.StructName}}Update true "更新{{.StructName}}"
// @Success 200 {object} response.Response{} "更新成功!"
// @Router /{{.Abbreviation}}/update [put]
func (s *{{.Abbreviation}}) Update(r *ghttp.Request) *response.Response {
	var info request.{{.StructName}}Update
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := example.{{.StructName}}.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags Example{{.StructName}}
// @Summary 删除 {{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /{{.Abbreviation}}/delete [delete]
func (s *{{.Abbreviation}}) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := example.{{.StructName}}.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// Deletes
// @Tags Example{{.StructName}}
// @Summary 批量删除 {{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByIDs true "批量删除{{.StructName}}"
// @Success 200 {object} response.Response{} "批量删除成功!"
// @Router /{{.Abbreviation}}/deletes [delete]
func (s *{{.Abbreviation}}) Delete{{.StructName}}ByIds(c *gin.Context) {
	var info common.GetByIDs
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := example.{{.StructName}}.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// GetList
// @Tags Example{{.StructName}}
// @Summary 分页获取 {{.Description}} 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.{{.StructName}}Search true "请求参数"
// @Success 200 {object} response.Response{data=[]example.{{.StructName}}} "获取列表数据成功!"
// @Router /{{.Abbreviation}}/getList [get]
func (s *{{.Abbreviation}}) GetList(c *gin.Context) {
    var info request.{{.StructName}}Search
    if err := r.Parse(&info); err != nil {
        return &response.Response{Error: err, MessageCode: response.ErrorGetList}
    }
    list, total, err := example.{{.StructName}}.GetList(&info)
    if err != nil {
        return &response.Response{Error: err, MessageCode: response.ErrorGetList}
    }
    return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}
