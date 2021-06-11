package service

import (
	"errors"
	model "flipped-aurora/gf-vue-admin/server/app/model/extra"
	"flipped-aurora/gf-vue-admin/server/app/model/extra/request"
	"flipped-aurora/gf-vue-admin/server/library/upload"
	"github.com/gogf/gf/frame/g"
	"mime/multipart"
	"strings"
)

var File = new(file)

type file struct{
	_file model.File
}

// Create
// Author: [SliverHorn](https://github.com/SliverHorn)
func (f *file) Create(info *request.CreateFile) (result *model.File, err error) {
	entity := info.Create()
	_, err = g.DB().Table(f._file.TableName()).Insert(entity)
	return entity, err
}

// First 根据id获取文件切片记录
// Author: [SliverHorn](https://github.com/SliverHorn)
func (f *file) First(info *request.GetById) (result *model.File, err error) {
	var entity model.File
	err = g.DB().Table(f._file.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

// Delete 删除文件记录
// Author: [SliverHorn](https://github.com/SliverHorn)
func (f *file) Delete(info *request.GetById) error {
	if result, err := f.First(info); err != nil {
		return err
	} else {
		if err = upload.Oss().Delete(result.Key); err != nil {
			return errors.New("文件删除失败")
		}
		_, err = g.DB().Table(f._file.TableName()).Delete(info.Condition())
		return err
	}
}

// GetList 分页获取数据
// Author: [SliverHorn](https://github.com/SliverHorn)
func (f *file) GetList(info *request.PageInfo) (list *[]model.File, total int, err error) {
	var files []model.File
	db := g.DB().Table(f._file.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&files)
	return &files, total, err
}

// UploadFile 根据配置文件判断是文件上传到本地或者七牛云
// Author: [SliverHorn](https://github.com/SliverHorn)
func (f *file) UploadFile(header *multipart.FileHeader, noSave string) (file *model.File, err error) {
	if filePath, key, uploadErr := upload.Oss().Upload(header); uploadErr != nil {
		return nil, uploadErr
	} else {
		s := strings.Split(header.Filename, ".")
		info := &request.CreateFile{BaseFile: request.BaseFile{Url: filePath, Tag: s[len(s)-1], Key: key, Name: header.Filename}}
		if noSave == "0" {
			return f.Create(info)
		}
		return nil, errors.New("不保存参数")
	}
}
