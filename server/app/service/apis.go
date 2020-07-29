package service

import (
	"errors"
	"server/app/api/request"
	"server/app/model/apis"

	"github.com/gogf/gf/frame/g"
)

// CreateApi 创建api
func CreateApi(api *request.CreateApi) error {
	_, err := apis.FindOne(g.Map{"path": api.Path, "method": api.Method})
	if err != nil {
		return errors.New("存在相同api")
	}
	apiToDb := &apis.Entity{
		Path:        api.Path,
		Description: api.Description,
		ApiGroup:    api.ApiGroup,
		Method:      api.Method,
	}
	if _, err := apis.Insert(apiToDb); err != nil {
		return errors.New("创建api失败")
	}
	return nil
}

// UpdateApi 更新api信息
func UpdateApi(api *request.UpdateApi) error {
	condition := g.Map{"id": api.ID}
	updateData := g.Map{
		"path":        api.Path,
		"description": api.Description,
		"api_group":   api.ApiGroup,
		"method":      api.Method,
	}

	oldA, err := apis.FindOne(condition)
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !apis.RecordNotFound(g.Map{"path": api.Path, "method": api.Method}) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	}
	if err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil {
		return err
	}
	_, err = apis.Update(updateData, condition)
	return err
}

// DeleteApi 删除api
func DeleteApi(api *request.DeleteApi) error {
	_, err := apis.Delete(g.Map{"id": api.ID})
	ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetApiById 根据id获取api
func GetApiById(api *request.GetById) (apisReturn *apis.Apis, err error) {
	apisReturn = (*apis.Apis)(nil)
	db := g.DB("default").Table("apis").Safe()
	err = db.Where(g.Map{"id": api.Id}).Struct(&apisReturn)
	return apisReturn, err
}

// GetAllApis 获取所有的Api
func GetAllApis() (list []*apis.Apis, err error) {
	list = ([]*apis.Apis)(nil)
	db := g.DB("default").Table("apis").Safe()
	err = db.Structs(&list)
	return list, err
}

// GetApiInfoList Page to get the data
// GetApiInfoList 分页获取数据
func GetApiInfoList(api *request.GetApiList) (list []*apis.Apis, total int, err error) {
	list = ([]*apis.Apis)(nil)
	db := g.DB("default").Table("apis").Safe()
	limit := api.PageSize
	offset := api.PageSize * (api.Page - 1)
	condition := g.Map{}
	if api.Path != "" {
		condition["path like ?"] = "%" + api.Path + "%"
	}
	if api.Description != "" {
		condition["description like ?"] = "%" + api.Description + "%"
	}
	if api.Method != "" {
		condition["method"] = api.Method
	}
	if api.ApiGroup != "" {
		condition["api_group"] = api.ApiGroup
	}
	total, err = db.Where(condition).Count()
	if api.OrderKey != "" && api.Desc == true {
		orderStr := api.OrderKey + " desc"
		err = db.Limit(limit).Offset(offset).Order(orderStr).Where(condition).Structs(&list)
		return list, total, err
	}
	if api.OrderKey != "" && api.Desc == false {
		err = db.Where(condition).Order("api_group").Limit(limit).Offset(offset).Structs(&list)
		return list, total, err
	}
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&list)
	return list, total, err
}
