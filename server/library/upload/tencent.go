package upload

import (
	"context"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"flipped-aurora/gf-vue-admin/server/library/upload/internal"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"mime/multipart"
	"time"
)

var TencentCos = new(tencent)

type tencent struct{
	err error
	file multipart.File
}

// Upload upload file to COS
// Author: [SliverHorn](https://github.com/SliverHorn)
func (t *tencent) Upload(file *multipart.FileHeader) (string, string, error) {
	client := internal.Tencent.Client()

	if t.file, t.err = file.Open(); t.err != nil {
		g.Log().Error("function file.Open() Failed!", g.Map{"err": t.err})
		return "", "", t.err
	}
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	if _, err := client.Object.Put(context.Background(), global.Config.Tencent.PathPrefix+"/"+fileKey, t.file, nil); err != nil {
		return "", "", err
	} else {
		return global.Config.Tencent.BaseURL + "/" + global.Config.Tencent.PathPrefix + "/" + fileKey, fileKey, nil
	}

}

// Delete delete file form COS
// Author: [SliverHorn](https://github.com/SliverHorn)
func (t *tencent) Delete(key string) error {
	client := internal.Tencent.Client()
	name := global.Config.Tencent.PathPrefix + "/" + key

	if response, err := client.Object.Delete(context.Background(), name); err != nil {
		g.Log().Error("function bucketManager.Delete() Failed!", g.Map{"err": err})
		return err
	} else {
		g.Log().Info("client.Object.Delete Success!", g.Map{"response": response})
		return nil
	}
}
