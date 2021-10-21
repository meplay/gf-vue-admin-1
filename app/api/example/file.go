package example

import (
	model "github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/service/example"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var File = new(file)

type file struct{}

// Upload
// @Tags ExampleFile
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "请求参数"
// @Success 200 {object} response.Response{data=example.File} "上传成功!"
// @Router /fileUploadAndDownload/upload [post]
func (a *file) Upload(r *ghttp.Request) *response.Response {
	noSave := r.GetQueryString("noSave", "0")
	_, header, err := r.Request.FormFile("file")
	if err != nil {
		return &response.Response{Error: err, Message: "接收文件失败!"}
	}
	var data *model.File
	data, err = example.File.Upload(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		return &response.Response{Error: err, Message: "接收文件失败!"}
	}
	return &response.Response{Data: g.Map{"file": data}, Message: "上传成功!"}
}

// Delete
// @Tags ExampleFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"\
// @Router /fileUploadAndDownload/deleteFile [post]
func (a *file) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := example.File.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags ExampleFile
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.PageInfo true "请求参数"
// @Success 200 {object} response.Response{data=[]example.File} "上传成功!"
// @Router /fileUploadAndDownload/getFileList [post]
func (a *file) GetList(r *ghttp.Request) *response.Response {
	var info common.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := example.File.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info), MessageCode: response.SuccessGetList}
}
