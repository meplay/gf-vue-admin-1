package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Menu = new(menu)

type menu struct{}

// Create
// @Tags SystemMenu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MenuCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /menu/addBaseMenu [post]
func (a *menu) Create(r *ghttp.Request) *response.Response {
	var info request.MenuCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := system.Menu.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{Message: "添加成功!"}
}

// First
// @Tags SystemMenu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{data=system.Menu} "获取数据成功!"
// @Router /menu/getBaseMenuById [post]
func (a *menu) First(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	data, err := system.Menu.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"menu": data}, MessageCode: response.SuccessFirst}
}

// Update
// @Tags SystemMenu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MenuUpdate true "请求参数"
// @Success 200 {object} response.Response{} "更新成功!"
// @Router /menu/updateBaseMenu [post]
func (a *menu) Update(r *ghttp.Request) *response.Response {
	var info request.MenuUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.Menu.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags SystemMenu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /menu/deleteBaseMenu [post]
func (a *menu) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := system.Menu.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags SystemMenu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.PageInfo true "请求参数"
// @Success 200 {object} response.Response{data=[]system.Menu} "获取列表数据成功!"
// @Router /menu/getMenuList [post]
func (a *menu) GetList(r *ghttp.Request) *response.Response {
	var info common.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.Menu.GetList()
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info), MessageCode: response.SuccessGetList}
}

// GetTree
// @Tags SystemMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body common.Empty true "空"
// @Success 200 {object} response.Response{data=[]system.Menu} "获取列表数据成功!"
// @Router /menu/getBaseMenuTree [post]
func (a *menu) GetTree(r *ghttp.Request) *response.Response {
	data, err := system.Menu.GetTree()
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: g.Map{"menus": data}, MessageCode: response.SuccessGetList}
}

// AddMenuAuthority
// @Tags SystemMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MenuAddAuthority true "请求参数"
// @Success 200 {object} response.Response{} "添加成功!"
// @Router /menu/addMenuAuthority [post]
func (a *menu) AddMenuAuthority(r *ghttp.Request) *response.Response {
	var info request.MenuAddAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "添加失败!"}
	}
	if err := system.Menu.AddMenuAuthority(&info); err != nil {
		return &response.Response{Error: err, Message: "添加失败!"}
	}
	return &response.Response{Message: "添加成功!"}
}
