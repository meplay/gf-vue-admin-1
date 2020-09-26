package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/dictionaries"
	"server/app/model/dictionary_details"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"
)

// CreateDictionary create a Dictionary
// CreateDictionary 创建Dictionary
func CreateDictionary(create *request.CreateDictionary) (err error) {
	var findData *dictionaries.Entity
	if findData, err = dictionaries.FindOne(g.Map{"type": create.Type}); err != nil {
		return errors.New("创建Dictionary失败")
	}
	if findData != nil {
		return errors.New("存在相同的type，不允许创建")
	}
	insert := &dictionaries.Entity{
		Name:   create.Name,
		Type:   create.Type,
		Status: utils.BoolToInt(create.Status),
		Desc:   create.Name,
	}
	if _, err = dictionaries.Insert(insert); err != nil {
		return errors.New("创建Dictionary失败")
	}
	return err
}

// DeleteDictionary delete a Dictionary
// DeleteDictionary 删除Dictionary
func DeleteDictionary(delete *request.DeleteDictionary) (err error) {
	if _, err = dictionaries.Delete(g.Map{"id": delete.Id}); err != nil {
		return errors.New("删除Dictionary失败")
	}
	_, err = dictionary_details.Delete(g.Map{"dictionary_id": delete.Id})
	return err
}

// UpdateDictionary update Dictionary
// UpdateDictionary 更新 Dictionary
func UpdateDictionary(update *request.UpdateDictionary) (err error) {
	var dictFromDb *dictionaries.Entity
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
		"name":   update.Name,
		"type":   update.Type,
		"status": update.Status,
		"desc":   update.Desc,
	}
	if dictFromDb, err = dictionaries.FindOne(condition); err != nil {
		return errors.New("记录不存在,更新失败")
	}
	if dictFromDb.Type == update.Type || dictionaries.RecordNotFound(g.Map{"type": update.Type}) {
		_, err = dictionaries.Update(updateData, condition)
		return err
	}
	return errors.New("更新失败")
}

// FindDictionary Find a Dictionary with id
// FindDictionary 用id查询Dictionary
func FindDictionary(find *request.FindDictionary) (dictionary *dictionaries.DictionaryHasManyDetails, err error) {
	db := g.DB("default").Table("dictionaries").Safe()
	detailDb := g.DB("default").Table("dictionary_details").Safe()
	err = db.Where(g.Map{"id": find.Id}).Or(g.Map{"`type`": find.Type}).Struct(&dictionary)
	err = detailDb.Structs(&dictionary.DictionaryDetails, "dictionary_id", dictionary.Id)
	return dictionary, err
}

// GetDictionaryInfoList get Dictionary list by pagination
// GetDictionaryInfoList 通过分页获得Dictionary列表
func GetDictionaryInfoList(info *request.DictionaryInfoList, condition g.Map) (list interface{}, total int, err error) {
	var dictionaryList []*dictionaries.Dictionaries
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	if info.Name != "" {
		condition["`name` like ?"] = "%" + info.Name + "%"
	}
	if info.Type != "" {
		condition["`type` like ?"] = "%" + info.Type + "%"
	}
	if info.Desc != "" {
		condition["`desc` like ?"] = "%" + info.Desc + "%"
	}
	db := g.DB("default").Table("dictionaries").Safe()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&dictionaryList)
	return dictionaryList, total, err
}
