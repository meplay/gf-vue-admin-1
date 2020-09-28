package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"server/library/utils"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

var localPath string

func init() {
	localPath = g.Cfg("oss").GetString("local.LocalPath")
}

type Local struct{}

func (l Local) Upload(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(localPath, os.ModePerm)
	if mkdirErr != nil {
		g.Log().Errorf("err:%v", mkdirErr)
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := localPath + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		g.Log().Errorf("err:%v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
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
	return p, filename, nil
}

func (l Local) DeleteFile(key string) error {
	p := localPath + "/" + key
	if strings.Contains(p, localPath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
