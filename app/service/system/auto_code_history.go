package system

import (
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
	"time"
)

type autoCodeHistory struct{}

var AutoCodeHistory = new(autoCodeHistory)

func (s *autoCodeHistory) Repeat(structName string) bool {
	var count int64
	global.Db.Model(&system.AutoCodeHistory{}).Where("struct_name = ? and flag = 0", structName).Count(&count)
	return count > 0
}

// Create RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func (s *autoCodeHistory) Create(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string) error {
	return global.Db.Create(&system.AutoCodeHistory{
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
		Request:       meta,
	}).Error
}

// RollBack 回滚
func (s *autoCodeHistory) RollBack(id uint) error {
	entity := system.AutoCodeHistory{}
	if err := global.Db.First(&entity, id).Error; err != nil {
		return err
	}

	err := Api.Deletes(entity.Apis.ToCommonGetByID())
	if err != nil {
		return errors.Wrap(err, "回滚api表数据失败!")
	} // 清除API表

	dbNames, getErr := AutoCode.GetTables(global.Config.Gorm.Dsn.Sources[0].GetDsn(global.Config.Gorm.Config))
	if getErr != nil {
		return errors.Wrap(getErr, "获取表数据失败!")
	} // 获取全部表名

	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.TableName, "_", "", -1)), strings.ToUpper(entity.TableName)) {
			if err = global.Db.Migrator().DropTable(name.TableName); err != nil {
				return errors.Wrap(err, "删除表失败!")
			} // 删除表
		}
	} // 删除表

	for _, path := range strings.Split(entity.AutoCodePath, ";") {

		_path, AbsErr := filepath.Abs(path) // 增加安全判断补丁
		if AbsErr != nil || _path != path {
			continue
		}

		// 迁移
		nPath := filepath.Join(global.Config.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		err = utils.File.Move(path, nPath)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>>", err)
		}
		//_ = utils.DeLFile(path)
	} // 删除文件

	for _, v := range strings.Split(entity.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.Injection.ClearAutoCode(meta[0], meta[2])
		}
	} // 清除注入
	entity.Flag = 1
	return global.Db.Save(&entity).Error
}

func (s *autoCodeHistory) GetMeta(id uint) (string, error) {
	var meta string
	return meta, global.Db.Model(system.AutoCodeHistory{}).Select("request_meta").First(&meta, id).Error
}

// GetList 获取系统历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) GetList(info *request.AutoCodeHistorySearch) (list []system.AutoCodeHistory, total int64, err error) {
	db := global.Db.Model(&system.AutoCodeHistory{})
	entities := make([]system.AutoCodeHistory, 0, info.PageSize)
	db = db.Scopes(info.Search())
	err = db.Count(&total).Scopes(common.Paginate(info.PageInfo), info.Order()).Find(&entities).Error
	return entities, total, err
}

// Delete 删除历史数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeHistory) Delete(id uint) error {
	return global.Db.Delete(system.AutoCodeHistory{}, id).Error
}
