package internal

import (
	"flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Tencent = new(tencent)

type tencent struct{}

// Client init COS client
func (t *tencent) Client() *cos.Client {
	_url, _ := url.Parse("https://" + global.Config.Tencent.Bucket + ".cos." + global.Config.Tencent.Region + ".myqcloud.com")
	client := cos.NewClient(&cos.BaseURL{BucketURL: _url}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.Config.Tencent.SecretID,
			SecretKey: global.Config.Tencent.SecretKey,
		},
	})
	return client
}
