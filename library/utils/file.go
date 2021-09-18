package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var File = new(file)

type file struct{}

// TrimSpace 去除结构体空格(目标结构体,传入必须是指针类型)
// Author [SliverHorn](https://github.com/SliverHorn)
func (f *file) TrimSpace(target interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	v := reflect.ValueOf(target).Elem()
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		}
	}
}

// Move 文件移动 src: 源位置,绝对路径or相对路径, dst: 目标位置,绝对路径or相对路径,必须为文件夹
// Author [SliverHorn](https://github.com/SliverHorn)
func (f *file) Move(src string, dst string) (err error) {
	if dst == "" || src == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	var revoke bool
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}
// ZipFiles
// Author [SliverHorn](https://github.com/SliverHorn)
func (f *file)ZipFiles(filename string, files []string, oldForm, newForm string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	for _, file := range files {
		err = func(file string) error {
			zipFile, openErr := os.Open(file)
			if openErr != nil {
				return openErr
			}
			defer func() { _ = zipFile.Close() }()

			info, statErr := zipFile.Stat() // 获取file的基础信息
			if statErr != nil {
				return statErr
			}

			header, HeaderErr := zip.FileInfoHeader(info)
			if HeaderErr != nil {
				return HeaderErr
			}

			// 使用上面的FileInfoHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
			header.Name = strings.Replace(file, oldForm, newForm, -1)

			// 优化压缩
			// 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
			header.Method = zip.Deflate

			writer, CreateErr := zipWriter.CreateHeader(header)
			if CreateErr != nil {
				return CreateErr
			}
			if _, err = io.Copy(writer, zipFile); err != nil {
				return err
			}
			return nil
		}(file)
		if err != nil {
			return err
		}
	} // 把files添加到zip中
	return nil
}