package oss

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var _ interfaces.Oss = (*local)(nil)

var Local = new(local)

type local struct {
	filename string
	filesize int64
}

func (l *local) DeleteByKey(key string) error {
	path := global.Config.Local.Path + "/" + key
	if strings.Contains(path, global.Config.Local.Path) {
		if err := os.Remove(path); err != nil {
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

func (l *local) UploadByFilepath(p string) (path string, filename string, err error) {
	var file *os.File
	file, err = os.Open(p)
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件不存在!", p)
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件信息获取失败!", p)
	}
	l.filesize = info.Size()
	_, l.filename = filepath.Split(path)
	return l.UploadByFile(file)
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
