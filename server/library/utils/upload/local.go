package upload

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"server/library/global"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

type Local struct{}

func (l Local) Upload(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	formatName, filenameErr := filename(file.Filename)
	if filenameErr != nil {
		g.Log().Errorf("err:%v", filenameErr)
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + filenameErr.Error())
	}
	// 拼接新文件名
	filename := formatName + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.Config.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		g.Log().Errorf("err:%v", mkdirErr)
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	path := global.Config.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		g.Log().Errorf("err:%v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(path)
	if createErr != nil {
		g.Log().Errorf("err:%v", createErr)
		return "", "", errors.New("function file.Open() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		g.Log().Errorf("err:%v", copyErr)
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return path, filename, nil
}

func (l Local) DeleteFile(url string) error {
	if strings.Contains(url, global.Config.Local.Path) {
		if err := os.Remove(url); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

func filename(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_sha256 := sha256.New()
	if _, err := io.Copy(_sha256, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(_sha256.Sum(nil)), nil
}
