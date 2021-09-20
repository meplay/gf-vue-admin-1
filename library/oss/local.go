package oss

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

var _ interfaces.Oss = (*local)(nil)

var Local = new(local)

type local struct {
	filename string
}

func (l *local) DeleteByKey(key string) error {
	filepath := global.Config.Local.Path + "/" + key
	if strings.Contains(filepath, global.Config.Local.Path) {
		if err := os.Remove(filepath); err != nil {
			return errors.Wrap(err, "本地文件删除失败!")
		}
	}
	return nil
}

func (l *local) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	filepath = global.Config.Local.Filepath(l.filename)
	var out *os.File
	if out, err = os.Create(filepath); err != nil {
		return filepath, filename, errors.Wrap(err, "读取文件失败!")
	}

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("open 文件关闭失败!", zap.Error(err))
		}
		if err = out.Close(); err != nil {
			zap.L().Error("out 文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	if _, err = io.Copy(out, file); err != nil {
		return filepath, filename, errors.Wrap(err, "传输(拷贝)文件失败!")
	} // 传输(拷贝)文件
	return filepath, filename, nil
}

func (l *local) UploadByFileHeader(file *multipart.FileHeader) (filepath string, filename string, err error) {
	if err = os.MkdirAll(global.Config.Local.Path, os.ModePerm); err != nil {
		return filepath, filename, errors.Wrap(err, "创建路径失败!")
	} // 尝试创建此路径
	l.filename = global.Config.Local.Filename(file.Filename)
	var open multipart.File
	if open, err = file.Open(); err != nil {
		return filepath, filename, errors.Wrap(err, "读取文件失败!")
	} // 读取文件
	return l.UploadByFile(open)
}
