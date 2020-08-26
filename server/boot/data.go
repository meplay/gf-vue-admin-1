package boot

import (
	"server/app/model"

	"github.com/gogf/gf/frame/g"
)

func InitializeDataTableAndData() {
	var err error
	// 初始化表,如果出错则不会执行初始化数据
	err = model.TableAdmins()
	err = model.TableApis()
	err = model.TableApis()
	err = model.TableAuthorities()
	err = model.TableAuthorityMenu()
	err = model.TableAuthorityResources()
	err = model.TableBreakpointChucks()
	err = model.TableBreakpointFiles()
	err = model.TableCasbinRule()
	err = model.TableCustomers()
	err = model.TableDictionaries()
	err = model.TableDictionaryDetails()
	err = model.TableFiles()
	err = model.TableJwts()
	err = model.TableMenus()
	err = model.TableOperations()
	if err != nil {
		g.Log().Error(err)
		return
	}
	// 初始化数据,并且数据插入是10条10条这样插入的,每个表插入数据都有加事务
	err = model.DataAdmins()
	err = model.DataApis()
	err = model.DataAuthorityMenus()
	err = model.DataCasbinRule()
	err = model.DataAuthorities()
	err = model.DataDictionaries()
	err = model.DataDictionaryDetails()
	err = model.DataMenus()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}
