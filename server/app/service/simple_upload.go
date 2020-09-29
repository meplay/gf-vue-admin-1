package service

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"server/app/api/request"
	"server/app/model/simple_upload"
	"server/library/global"
	"server/library/utils"
	"strconv"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
)

// 保存文件切片路径
func SaveChunk(uploader *request.CreateSimpleUpload) (err error) {
	insert := simple_upload.Entity{
		ChunkNumber:      uploader.ChunkNumber,
		CurrentChunkPath: uploader.CurrentChunkPath,
		CurrentChunkSize: uploader.CurrentChunkSize,
		Filename:         uploader.Filename,
		Identifier:       uploader.Identifier,
		TotalChunks:      uploader.TotalChunks,
		TotalSize:        uploader.TotalSize,
	}
	_, err = simple_upload.Insert(&insert)
	return err
}

// 检查文件是否已经上传过
func CheckFileMd5(md5 string) (uploads []*simple_upload.SimpleUpload, isDone bool, err error) {
	uploads = ([]*simple_upload.SimpleUpload)(nil)
	db := g.DB(global.Db).Table("simple_upload").Safe()
	err = db.Where(g.Map{"identifier": md5, "is_done": false}).Structs(&uploads)
	isDone = simple_upload.RecordNotFound(g.Map{"identifier": md5, "is_done": true})
	return uploads, !isDone, err
}

// MergeFileMd5 Merge File With Md5
// MergeFileMd5 合并文件
func MergeFileMd5(md5 string, fileName string) (err error) {
	finishDir := "./finish/"
	dir := "./chunk/" + md5
	if !simple_upload.RecordNotFound(g.Map{"identifier": md5, "is_done": true}) { // 如果文件上传成功 不做后续操作 通知成功即可
		return nil
	}
	//打开切片文件夹
	rd, readErr := ioutil.ReadDir(dir)
	if readErr != nil {
		return readErr
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	//创建目标文件
	fd, _ := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}
	defer fd.Close() //关闭文件

	var tx *gdb.TX
	if tx, err = g.DB("default").Begin(); err != nil { //创建事务
		panic(err)
	}
	db := g.DB(global.Db).Table("simple_upload").Safe()
	if _, err := db.Delete(g.Map{"identifier": md5, "is_done": false}); err != nil { // 删除切片信息
		return tx.Rollback()
	}
	// 添加文件信息
	insert := &simple_upload.Entity{
		FilePath:   finishDir + fileName,
		Filename:   fileName,
		Identifier: md5,
		IsDone:     utils.BoolToInt(true),
	}
	if _, err = db.Insert(insert); err != nil {
		return tx.Rollback()
	}
	err = os.RemoveAll(dir) //清除切片
	return tx.Commit()
}

func Upload(file *multipart.FileHeader, chunkPath string) error {
	f, openError := file.Open() // 读取文件
	if openError != nil {
		g.Log().Errorf("err:%v", openError)
		return errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(chunkPath)
	if createErr != nil {
		g.Log().Errorf("err:%v", createErr)
		return errors.New("function file.Open() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		g.Log().Errorf("err:%v", copyErr)
		return errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return nil
}
