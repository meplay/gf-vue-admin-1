package v1

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"server/app/model/breakpoint_files"
	"server/app/service"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// BreakpointContinue The breakpoint continues to the server
// BreakpointContinue 断点续传到服务器
func BreakpointContinue(r *ghttp.Request) {
	var (
		pathc string
		f     multipart.File
		file  *breakpoint_files.Entity
	)
	fileMd5 := r.Request.FormValue("fileMd5")
	fileName := r.Request.FormValue("fileName")
	chunkMd5 := r.Request.FormValue("chunkMd5")
	chunkNumber := gconv.Int(r.Request.FormValue("chunkNumber"))
	chunkTotal := gconv.Int(r.Request.FormValue("chunkTotal"))
	_, fileHeader, err := r.Request.FormFile("file")
	if err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if f, err = fileHeader.Open(); err != nil {
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
	if err = service.CreateFileChunk(file.Id, pathc, chunkNumber); err != nil {
		global.FailWithMessage(r, err.Error())
		return
	}
	global.OkWithMessage(r, "切片创建成功")
}

// FindFile Find files
// FindFile 查找文件
func FindFile(r *ghttp.Request) {
	fileMd5 := gconv.String(r.GetQuery("fileMd5"))
	fileName := gconv.String(r.GetQuery("fileName"))
	chunkTotal := gconv.Int(r.GetQuery("chunkTotal"))
	file, err := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.FailWithMessage(r, "查找失败")
		r.Exit()
	}
	global.OkWithData(r, g.Map{"file": file})
}

// BreakpointContinueFinish Find files
// BreakpointContinueFinish 查找文件
func BreakpointContinueFinish(r *ghttp.Request) {
	fileMd5 := r.GetString("fileMd5")
	fileName := r.GetString("fileName")
	filePath, err := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.FailWithDetailed(r, global.ERROR, g.Map{"filePath": filePath}, fmt.Sprintf("文件创建失败：%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"filePath": filePath}, "文件创建成功")

}

// RemoveChunk Delete slices
// RemoveChunk 删除切片
func RemoveChunk(r *ghttp.Request) {
	fileMd5 := r.GetString("fileMd5")
	fileName := r.GetString("fileName")
	filePath := r.GetString("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.FailWithDetailed(r, global.ERROR, g.Map{"filePath": filePath}, fmt.Sprintf("缓存切片删除失败：%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"filePath": filePath}, "缓存切片删除成功")
}
