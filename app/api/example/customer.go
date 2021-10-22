package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/example"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Customer = new(customer)

type customer struct {}

// Create
// @Tags ExampleCustomer
// @Summary 创建 客户表
// @Security ApiKeyAuth
// @accept application/json
// @Param data body request.CustomerCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /customer/createCustomer [post]

func (s *customer) Create(r *ghttp.Request) *response.Response{
	var info request.CustomerCreate
	if err := r.Parse(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := example.Customer.Create(&info);err != nil{
		return &response.Response{Error: err,MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}



// First
// @Tags ExampleCustomer
// @Summary 用id查询 客户表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "获取数据成功!"
// @Router /customer/first [get]
func (s *customer) First(r *ghttp.Request) *response.Response{
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	data,err := example.Customer.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"customer": data}, MessageCode: response.SuccessFirst}
}

// Update
// @Tags ExampleCustomer
// @Summary 更新 客户表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CustomerUpdate true "更新Customer"
// @Success 200 {object} response.Response{} "更新成功!"
// @Router /customer/update [put]
func (s *customer) Update(r *ghttp.Request) *response.Response{
	var info request.CustomerUpdate
	if err := r.Parse(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := example.Customer.Update(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags ExampleCustomer
// @Summary 删除 客户表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /customer/delete [delete]
func (s *customer) Delete(r *ghttp.Request) *response.Response{
	var info common.GetByID
	if err := r.Parse(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := example.Customer.Delete(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// Deletes
// @Tags ExampleCustomer
// @Summary 批量删除 客户表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByIDs true "批量删除Customer"
// @Success 200 {object} response.Response{} "批量删除成功!"
// @Router /customer/deletes [delete]
func (s *customer) Deletes(r *ghttp.Request) *response.Response{
	var info common.GetByIDs
	if err := r.Parse(&info); err !=nil{
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := example.Customer.Deletes(&info);err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// GetList
// @Tags ExampleCustomer
// @Summary 分页获取 客户表 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.CustomerSearch true "请求参数"
// @Success 200 {object} response.Response{data=[]example.Customer} "获取列表数据成功!"
// @Router /customer/getList [get]
func (s *customer) GetList(r *ghttp.Request) *response.Response{
	var info request.CustomerSearch
	if err := r.Parse(&info); err != nil{
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	list, total, err := example.Customer.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}