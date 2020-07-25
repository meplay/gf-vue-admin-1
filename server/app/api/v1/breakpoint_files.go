package v1

import (
	"io/ioutil"
	"mime/multipart"
	"server/app/service"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func BreakpointContinue(r *ghttp.Request) {
	var (
		pathc string
		f     multipart.File
	)
	fileMd5 := r.Request.FormValue("fileMd5")
	fileName := r.Request.FormValue("fileName")
	chunkMd5 := r.Request.FormValue("chunkMd5")
	chunkNumber := gconv.Int(r.Request.FormValue("chunkNumber"))
	chunkTotal := gconv.Int(r.Request.FormValue("chunkTotal"))
	_, FileHeader, err := r.Request.FormFile("file")
	if err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if f, err = FileHeader.Open(); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if flag := utils.CheckMd5(cen, chunkMd5); !flag {
		return
	}
	if file, err = service.FindOrCreateFile(fileMd5, fileName, chunkTotal); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	pathc, err = utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err = service.CreateFileChunk(bk_files.Id, pathc, chunkNumber); err != nil {
		global.FailWithMessage(r, err.Error())
		return
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func FindFile(r *ghttp.Request) {
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func BreakpointContinueFinish(r *ghttp.Request) {
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(r *ghttp.Request) {
	//fileMd5 := r.GetString("fileMd5")
	//fileName := r.GetString("fileName")
	//filePath := r.GetString("filePath")
	//err := utils.RemoveChunk(fileMd5)
	//err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	//if err != nil {
	//	global.FailWithDetailed(r, global.ERROR, g.Map{"filePath": filePath}, fmt.Sprintf("缓存切片删除失败：%v", err))
	//	r.Exit()
	//}
	//global.OkDetailed(r, g.Map{"filePath": filePath}, "缓存切片删除成功")
}
