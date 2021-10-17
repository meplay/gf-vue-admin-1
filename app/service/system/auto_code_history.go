package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
)

var AutoCodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// Create 创建代码生成器历史记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) Create(info *request.AutoCodeHistoryCreate) error {
	entity := info.Create()
	return global.Db.Create(&entity).Error
}

// First 根据id寻找代码生成器历史记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) First(info *common.GetByID) (*system.AutoCodeStruct, error) {
	var entity system.AutoCodeHistory
	err := global.Db.Model(system.AutoCodeHistory{}).Select("request").First(&entity, info.ID).Error
	if err != nil {
		return nil, errors.Wrap(err, "代码生成器历史记录寻找失败!")
	}
	return &entity.Request, err
}

// Repeat 判断重复
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) Repeat(structName string) bool {
	var count int64
	global.Db.Model(&system.AutoCodeHistory{}).Where("struct_name = ? and flag = 0", structName).Count(&count)
	return count > 0
}

// RollBack 回滚代码生成器历史记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) RollBack(info *common.GetByID) error {
	var entity system.AutoCodeHistory
	if err := global.Db.First(&entity, info.ID).Error; err != nil {
		return err
	}

	err := Api.Deletes(entity.Apis.ToCommonGetByID())
	if err != nil {
		return errors.Wrap(err, "回滚api表数据失败!")
	} // 清除API表

	var dbNames []response.Table
	dbNames, err = AutoCode.GetTables(global.Config.Gorm.Dsn.GetDefaultDsn(global.Config.Gorm.Config))
	if err != nil {
		return errors.Wrap(err, "获取表数据失败!")
	} // 获取全部表名

	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.TableName, "_", "", -1)), strings.ToUpper(entity.TablesName)) {
			if err = global.Db.Migrator().DropTable(name.TableName); err != nil {
				return errors.Wrap(err, "删除表失败!")
			} // 删除表
		} // 寻找相同表名
	} // 删除代码生成器生成的表

	for _, path := range entity.AutoCodePaths {
		var _path string
		_path, err = filepath.Abs(path.Filepath) // 增加安全判断补丁
		if err != nil || _path != path.Filepath {
			continue
		}
		newPath := path.RmFilePath()                  // 生成迁移路径
		err = utils.File.Move(path.Filepath, newPath) // 迁移
		if err != nil {
			return errors.Wrap(err, "迁移失败!")
		}
		//_ = utils.DeLFile(path)
	} // 删除文件

	for _, injection := range entity.Injection {
		if err = utils.Injection.ClearAutoCode(injection.Filepath, injection.CodeData); err != nil {
			return errors.Wrap(err, "清除注入代码失败!")
		}
	} // 清除注入
	return global.Db.Model(&system.AutoCodeHistory{}).Where("id =?", info.ID).Update("flag", 1).Error
}

// Delete 删除历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) Delete(info *common.GetByID) error {
	return global.Db.Delete(&system.AutoCodeHistory{}, info.ID).Error
}

// GetList 获取系统历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) GetList(info *request.AutoCodeHistorySearch) (list []system.AutoCodeHistory, total int64, err error) {
	db := global.Db.Model(&system.AutoCodeHistory{})
	entities := make([]system.AutoCodeHistory, 0, info.PageSize)
	err = db.Count(&total).Scopes(info.Select(), common.Paginate(info.PageInfo), info.Order()).Find(&entities).Error
	return entities, total, err
}
