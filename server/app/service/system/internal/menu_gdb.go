package internal

import (
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Menu = new(menu)

type menu struct {
	_parameters      model.MenuParameter
	_menusParameters model.MenusParameters
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取菜单的Parameters
func (m *menu) GetMenusParameters(id uint) *[]model.MenuParameter {
	entities := make([]model.MenusParameters, 0, 10)
	if err := g.DB().Table(m._menusParameters.TableName()).Where(g.Map{"base_menu_id": id}).Structs(&entities); err != nil {
		g.Log().Error("获取menus_parameters表数据失败!", g.Map{"base_menu_id": id})
		return nil
	}
	parameters := make([]model.MenuParameter, 0, len(entities))
	for _, entity := range entities{
		var e1 model.MenuParameter
		if err := g.DB().Table(m._parameters.TableName()).Where(g.Map{"parameter_id": entity.ParameterId}).Struct(&e1); err != nil {
			g.Log().Error("获取menus_parameters表数据失败", g.Map{"parameter_id": entity.ParameterId})
		} else {
			parameters = append(parameters, e1)
		}
	}
	return &parameters
}

func (m *menu) GetAuthoritiesMenus()  {

}