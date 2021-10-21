package example

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type File struct {
	global.Model
	Key  string `json:"key" gorm:"comment:编号"`
	Url  string `json:"url" gorm:"comment:文件地址"`
	Tag  string `json:"tag" gorm:"comment:文件标签"`
	Name string `json:"name" gorm:"comment:文件名"`
}
