package example

import (
	"mime/multipart"
	"strings"

	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/oss"
	"github.com/pkg/errors"
)

var File = new(file)

type file struct{}

// Upload 上传文件
// Author SliverHorn
func (s *file) Upload(header *multipart.FileHeader, noSave string) (*example.File, error) {
	filepath, key, err := oss.Oss().UploadByFileHeader(header)
	if err != nil {
		return nil, err
	}
	if noSave == "0" {
		slice := strings.Split(header.Filename, ".")
		info := request.FileCreate{
			Url:  filepath,
			Name: header.Filename,
			Tag:  slice[len(slice)-1],
			Key:  key,
		}
		return s.Create(&info)
	}
	return nil, nil
}

// Create 创建文件上传记录
// Author SliverHorn
func (s *file) Create(info *request.FileCreate) (*example.File, error) {
	entity := info.Create()
	err := global.Db.Create(&entity).Error
	return &entity, err
}

// First 根据id获取文件记录
// Author SliverHorn
func (s *file) First(info *common.GetByID) (*example.File, error) {
	var entity example.File
	err := global.Db.First(entity, info.ID).Error
	return &entity, err
}

// Delete 删除文件记录
// Author SliverHorn
func (s *file) Delete(info *common.GetByID) error {
	entity, err := s.First(info)
	if err != nil {
		return errors.Wrap(err, "文件记录不存在!")
	}
	if err = oss.Oss().DeleteByKey(entity.Key); err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}
	return global.Db.Unscoped().Delete(&entity).Error
}

// GetList 分页获取文件记录数据
// Author SliverHorn
func (s *file) GetList(info *common.PageInfo) (list []example.File, total int64, err error) {
	db := global.Db.Model(&example.File{})
	var entities []example.File
	err = db.Count(&total).Scopes(common.Paginate(info)).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}
