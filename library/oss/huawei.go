package oss

import (
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

var _ interfaces.OSS = (*_obs)(nil)

var HuaWeiObs = new(_obs)

type _obs struct{}

func (o *_obs) Delete(key string) error {
	client, err := obs.New(global.Config.HuaWeiObs.AccessKey, global.Config.HuaWeiObs.SecretKey, global.Config.HuaWeiObs.Endpoint)
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.Config.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	file, _ := os.Open("s")
	info, _ := file.Stat()
	o.Upload(file, NewFileInfoByFs(info, WithReader(file)))

	var header *multipart.FileHeader
	file2, _ := header.Open()
	info2 := NewFileInfo(
		WithFilesize(header.Size),
		WithFilename(header.Filename),
		WithFileMode(fs.ModePerm),
		WithFileModTime(time.Now().Local()),
		WithFileIsDir(false),
		WithFileSys(nil),
	)
	o.Upload(file2, info2)
	return nil
}

func (o *_obs) Upload(reader io.Reader, info interfaces.FileInfo) (filepath string, filename string, err error) {
	buf := make([]byte, 0)
	_, err = reader.Read(buf)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "读取文件[]byte失败!")
	}
	http.DetectContentType(buf)
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.Config.HuaWeiObs.Bucket,
				Key:    filename,
			},
			ContentType: http.DetectContentType(buf),
		},
		Body: reader,
	}

	var client *obs.ObsClient
	client, err = obs.New(global.Config.HuaWeiObs.AccessKey, global.Config.HuaWeiObs.SecretKey, global.Config.HuaWeiObs.Endpoint)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "文件上传失败!")
	}
	filepath = global.Config.HuaWeiObs.Path + "/" + filename
	return filepath, filename, err
}
