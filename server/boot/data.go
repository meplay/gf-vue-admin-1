package boot

import (
	"server/app/model"

	"github.com/gogf/gf/frame/g"
)

func InitializeDataTableAndData() {
	var err error
	// 初始化表,如果出错则不会执行初始化数据
	err = model.TableApis()
	err = model.TableJwts()
	err = model.TableMenus()
	err = model.TableFiles()
	err = model.TableAdmins()
	err = model.TableCustomers()
	err = model.TableCasbinRule()
	err = model.TableOperations()
	err = model.TableParameters()
	err = model.TableAuthorities()
	err = model.TableSimpleUpload()
	err = model.TableDictionaries()
	err = model.TableAuthorityMenu()
	err = model.TableBreakpointFiles()
	err = model.TableBreakpointChucks()
	err = model.TableDictionaryDetails()
	err = model.TableAuthorityResources()
	if err != nil {
		g.Log().Error(err)
		return
	}
	// 初始化数据,并且数据插入是10条10条这样插入的,每个表插入数据都有加事务
	err = model.DataApis()
	err = model.DataFiles()
	err = model.DataMenus()
	err = model.DataAdmins()
	err = model.DataCasbinRule()
	err = model.DataAuthorities()
	err = model.DataDictionaries()
	err = model.DataAuthorityMenus()
	err = model.DataDictionaryDetails()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}
