package utils

import (
	"errors"
	"fmt"
	"gf-vue-admin/interfaces"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var Excel = new(_excel)

type _excel struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 导出xlsx 通过面向接口方式, 数据由后端处理
func (e *_excel) ExportByInterface(export interfaces.ExcelExport) error {
	a1 := export.A1Data()
	data := export.DataList()
	name := export.SheetName()
	filepath := export.FilePath()
	if filepath == "" || len(a1) == 0 || len(data) == 0 {
		return errors.New("文件名 or A1Data 数据 or DataList数据 不能为空! ")
	}
	excel := excelize.NewFile()
	if name != "" {
		excel.SetSheetName("Sheet1", name)
	} else {
		name = "Sheet1"
	}
	if err := excel.SetSheetRow(name, "A1", &a1); err != nil {
		return err
	}
	for i, d := range data {
		index := fmt.Sprintf("A%d", i+2)
		if err := excel.SetSheetRow(name, index, &d); err != nil {
			return err
		}
	}
	return excel.SaveAs(filepath)
}
