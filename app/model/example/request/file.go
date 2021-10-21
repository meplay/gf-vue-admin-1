package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
)

type FileCreate struct {
	Key  string `json:"key" gorm:"comment:编号"`
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Name string `json:"name" gorm:"comment:文件名"`
}

func (r *FileCreate) Create() example.File {
	return example.File{
		Key:  r.Key,
		Url:  r.Url,
		Tag:  r.Tag,
		Name: r.Name,
	}
}
