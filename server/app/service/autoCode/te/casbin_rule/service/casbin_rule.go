package service

import (
	"server/app/api/request"
	"server/app/model/admins"
	"server/app/model/casbin_rule"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

// CreateCasbinRule create a CasbinRule
// CreateCasbinRule 创建CasbinRule
func CreateCasbinRule(create *request.CreateCasbinRule) (err error) {
	insert := casbin_rule.Entity{
        p_type: create.PType,
        v0: create.V0,
        v1: create.V1,
        v2: create.V2,
        v3: create.V3,
        v4: create.V4,
        v5: create.V5,
	}
	_, err = operations.Insert(&insert)
	return err
}

// DeleteCasbinRule delete CasbinRule
// DeleteCasbinRule 删除 CasbinRule
func DeleteCasbinRule(delete *request.DeleteById) (err error) {
	_, err = operations.Delete(g.Map{"id": delete.Id})
	return err
}

// DeleteCasbinRules batch deletion CasbinRules
// DeleteCasbinRules 批量删除 CasbinRules
func DeleteCasbinRules(deletes *request.DeleteByIds) (err error) {
	_, err = operations.Delete(g.Map{"id IN(?)": deletes.Ids})
	return err
}

// UpdateCasbinRule update CasbinRules
// UpdateCasbinRule 更新 CasbinRules
func UpdateCasbinRule(update *request.UpdateCasbinRule) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
        p_type: update.PType,
        v0: update.V0,
        v1: update.V1,
        v2: update.V2,
        v3: update.V3,
        v4: update.V4,
        v5: update.V5,
	}
	_, err = operations.Update(updateData, condition)
	return err
}

// FindCasbinRule Gets a single CasbinRule based on id
// FindCasbinRule 根据id获取单条CasbinRule
func FindCasbinRule(find *request.FindCasbinRule) (data *casbin_rule.Entity, err error) {
	return operations.FindOne(g.Map{"id": find.Id})
}

// GetCasbinRuleList Page out the CasbinRule list
// GetCasbinRuleList 分页获取CasbinRule列表
func GetCasbinRuleList(info *request.GetCasbinRuleList, condition g.Map) (list interface{}, total int, err error) {
	datalist := ([]*casbin_rule.CasbinRule)(nil)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB(global.Db).Table("casbin_rule").Safe()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Structs(&datalist)
	return datalist, total, err
}
