package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/service"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CreateSimpleUpload Example of resumable upload plug-in version
// CreateSimpleUpload 断点续传插件版示例
func CreateSimpleUpload(r *ghttp.Request) {
	var chunk request.CreateSimpleUpload
	_, header, err := r.Request.FormFile("file")
	chunk.Filename = r.PostFormValue("filename")
	chunk.ChunkNumber = r.PostFormValue("chunkNumber")
	chunk.CurrentChunkSize = r.PostFormValue("currentChunkSize")
	chunk.Identifier = r.PostFormValue("identifier")
	chunk.TotalSize = r.PostFormValue("totalSize")
	chunk.TotalChunks = r.PostFormValue("totalChunks")
	fmt.Println("Identifier", chunk.Identifier)
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		_ = utils.CreateDir(chunkDir)
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	if header != nil {
		err = service.Upload(header, chunkPath)
	}
	chunk.CurrentChunkPath = chunkPath
	err = service.SaveChunk(&chunk)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("切片创建失败，%v", err.Error()))
		return
	}
	global.Ok(r)
}

// CheckFileMd5 Test whether the file already exists and determine the uploaded slice
// CheckFileMd5 测试文件是否已经存在和判断已经上传过的切片
func CheckFileMd5(r *ghttp.Request) {
	md5 := r.GetQueryString("md5")
	chunks, isDone, err := service.CheckFileMd5(md5)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("md5读取失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"chunks": chunks, "isDone": isDone})
}

// MergeFileMd5 Merge files
// MergeFileMd5 合并文件
func MergeFileMd5(r *ghttp.Request) {
	md5 := r.GetQueryString("md5")
	fileName := r.GetQueryString("fileName")
	err := service.MergeFileMd5(md5, fileName)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("md5读取失败，%v", err))
	} else {
		global.OkWithData(r, g.Map{})
	}
}
