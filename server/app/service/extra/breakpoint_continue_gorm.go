package service

import (
	"errors"
	model "gf-vue-admin/app/model/extra"
	"gf-vue-admin/app/model/extra/request"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/utils"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"strconv"
)

var BreakpointContinueGorm = new(breakpointContinueGorm)

type breakpointContinueGorm struct {
	err     error
	path    string
	file    multipart.File
	_entity model.BreakpointContinue
	entity  *model.BreakpointContinue
}

// FindOrCreateFile 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *breakpointContinueGorm) FindOrCreateFile(info *request.BreakpointContinue) (result *model.BreakpointContinue, err error) {
	var entity model.BreakpointContinue
	var create = info.Create()
	if errors.Is(global.Db.Where("file_md5 = ? AND is_finish = ?", info.FileMd5, true).First(&entity).Error, gorm.ErrRecordNotFound) {
		err = global.Db.Where("file_md5 = ? AND file_name = ?", info.FileMd5, info.FileName).Preload("FileChunk").FirstOrCreate(&entity, create).Error
		return &entity, err
	}
	create.IsFinish = true
	create.FilePath = entity.FilePath
	err = global.Db.Create(&entity).Error
	return &entity, err
}

// CreateFileChunk 创建文件切片记录
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *breakpointContinueGorm) CreateFileChunk(info *request.CreateFileChunk) error {
	entity := info.Create()
	return global.Db.Create(&entity).Error
}

// DeleteFileChunk 删除文件切片记录
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *breakpointContinueGorm) DeleteFileChunk(info *request.BreakpointContinue) error {
	var chunks []model.BreakpointContinueChunk
	var entity model.BreakpointContinue
	var err = global.Db.Where("file_md5 = ? AND file_name = ?", info.FileMd5, info.FileName).First(&entity).Updates(map[string]interface{}{"is_finish": true, "file_path": info.FilePath}).Error
	err = global.Db.Where("file_id = ?", entity.ID).Delete(&chunks).Unscoped().Error
	err = utils.File.RemoveChunk(info.FileMd5)
	return err
}

// BreakpointContinue 断点续传到服务器
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *breakpointContinueGorm) BreakpointContinue(info *request.BreakpointContinue, header *multipart.FileHeader) error {
	if b.file, b.err = header.Open(); b.err != nil {
		return errors.New("文件读取失败! ")
	}
	content, _ := ioutil.ReadAll(b.file)
	if !utils.File.CheckMd5(content, info.ChunkMd5) {
		return errors.New("检查md5失败! ")
	}
	b.entity, b.err = b.FindOrCreateFile(info)
	chunkNumber, _ := strconv.Atoi(info.ChunkNumber)
	chunkTotal, _ := strconv.Atoi(info.ChunkTotal)
	if b.path, b.err = utils.File.BreakPointContinue(content, info.FileName, chunkNumber, chunkTotal, info.FileMd5); b.err != nil {
		return errors.New("断点续传失败! ")
	}
	create := request.CreateFileChunk{
		GetById:                     request.GetById{Id: b.entity.ID},
		BaseBreakpointContinueChunk: request.BaseBreakpointContinueChunk{FileChunkPath: b.path, FileChunkNumber: chunkNumber},
	}
	if b.err = b.CreateFileChunk(&create); b.err != nil {
		return errors.New("创建文件记录失败! ")
	}
	return nil
}

// BreakpointContinueFinish 上传文件完成
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *breakpointContinueGorm) BreakpointContinueFinish(info *request.BreakpointContinueFinish) (filepath string, err error) {
	return utils.File.MakeFile(info.FileName, info.FileMd5)
}
