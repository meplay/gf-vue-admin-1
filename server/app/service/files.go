package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/breakpoint_files"
	"server/app/model/files"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// UploadFile Create file upload records
// UploadFile 创建文件上传记录
func UploadFile(file *files.Entity) (err error) {
	_, err = files.Insert(file)
	return err
}

// DeleteFile Delete file records
// DeleteFile 删除文件记录
func DeleteFile(d *request.DeleteFile) error {
	if _, err := files.FindOne(g.Map{"id": d.Id}); err != nil {
		return errors.New("文件记录不存在,删除失败")
	}
	_, err := files.Delete(g.Map{"id": d.Id})
	return err
}

// GetFileList Paging fetch data
// GetFileList 分页获取数据
func GetFileList(info *request.PageInfo) (list interface{}, total int, err error) {
	var fileList []*files.Entity
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("files").Safe()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Scan(&fileList)
	return fileList, total, err
}

func FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file breakpoint_files.Entity, err error) {
	insert := g.Map{
		"file_md5":    fileMd5,
		"file_name":   fileName,
		"chunk_total": chunkTotal,
	}
	if breakpoint_files.RecordNotFound(g.Map{"file_md5": fileMd5, "is_finish": utils.BoolToInt(true)}) {

	}
	insert["is_finish"] = utils.BoolToInt(true)
	insert["file_path"] = file.FilePath
	_, err = breakpoint_files.Insert(insert)
	return
}

func CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {

}
