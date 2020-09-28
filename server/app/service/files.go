package service

import (
	"errors"
	"mime/multipart"
	"server/app/api/request"
	"server/app/model/breakpoint_chucks"
	"server/app/model/breakpoint_files"
	"server/app/model/files"
	"server/library/utils"
	"server/library/utils/upload"
	"strings"

	"github.com/gogf/gf/frame/g"
)

// CreateFile Create file upload records
// CreateFile 创建文件上传记录
func CreateFile(file *files.Entity) (err error) {
	_, err = files.Insert(file)
	return err
}

// FindFile Find file record
// FindFile 查找文件记录
func FindFile(key string) (file *files.Files, err error) {
	file = (*files.Files)(nil)
	db := g.DB("default").Table("files").Safe()
	err = db.Where(g.Map{"key": key}).Struct(&file)
	return file, err
}

// FindFile Find file record By Id
// FindFile 根据Id查找文件记录
func FindFileById(find *request.DeleteFile) (file *files.Files, err error) {
	file = (*files.Files)(nil)
	db := g.DB("default").Table("files").Safe()
	err = db.Where(g.Map{"id": find.Id}).Struct(&file)
	return file, err
}

// DeleteFile Delete file records
// DeleteFile 删除文件记录
func DeleteFile(d *request.DeleteFile) error {
	fileFromDb, err := FindFileById(d)
	err = upload.Oss.DeleteFile(fileFromDb.Key)
	if _, err := files.FindOne(g.Map{"id": d.Id}); err != nil {
		return errors.New("文件记录不存在,删除失败")
	}
	_, err = files.Delete(g.Map{"id": d.Id})
	return err
}

// GetFileList Paging fetch data
// GetFileList 分页获取数据
func GetFileList(info *request.PageInfo) (list interface{}, total int, err error) {
	fileList := ([]*files.Files)(nil)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB("default").Table("files").Safe()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Scan(&fileList)
	return fileList, total, err
}

// FindOrCreateFile Check your file if it does not exist, or return current slice of the file
// FindOrCreateFile 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
func FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file *breakpoint_files.Entity, err error) {
	insert := g.Map{
		"file_md5":    fileMd5,
		"file_name":   fileName,
		"chunk_total": chunkTotal,
	}
	if breakpoint_files.RecordNotFound(g.Map{"file_md5": fileMd5, "is_finish": utils.BoolToInt(true)}) {
		return breakpoint_files.FindOne(g.Map{"file_md5": fileMd5})
	}
	insert["is_finish"] = utils.BoolToInt(true)
	insert["file_path"] = file.FilePath // TODO file.FilePath
	_, err = breakpoint_files.Insert(insert)
	return breakpoint_files.FindOne(g.Map{"file_md5": fileMd5})
}

// CreateFileChunk create a chunk of the file
// CreateFileChunk 创建文件切片记录
func CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) (err error) {
	insert := g.Map{
		"exa_file_id":       id,
		"file_chunk_path":   fileChunkPath,
		"file_chunk_number": fileChunkNumber,
	}
	_, err = breakpoint_chucks.Insert(insert)
	return err
}

// DeleteFileChunk delete a chuck of the file
// DeleteFileChunk 删除文件切片记录
func DeleteFileChunk(fileMd5 string, fileName string, filePath string) (err error) {
	var files *breakpoint_files.Entity
	condition := g.Map{
		"file_md5":  fileMd5,
		"file_name": fileName,
	}
	_, err = breakpoint_files.Update(condition, g.Map{"is_finish": utils.BoolToInt(true)})
	files, err = breakpoint_files.FindOne(g.Map{"file_md5": fileMd5})
	_, err = breakpoint_chucks.Delete(g.Map{"exa_file_id": files.ChunkId})
	return err
}

// UploadFile 根据配置文件进行文件上传
func UploadFile(header *multipart.FileHeader, noSave string) (file *files.Files, err error) {
	p, key, uploadErr := upload.Oss.Upload(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := &files.Entity{
			Url:  p,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		if err := CreateFile(f); err != nil {
			return nil, err
		}
		return FindFile(f.Key)
	}
	return
}
