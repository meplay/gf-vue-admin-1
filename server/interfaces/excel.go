package interfaces

type ExcelExport interface {
	A1Data() []string          // 表单
	FilePath() string          // 文件保存路径
	DataList() [][]interface{} // 数据
	SheetName() string         // 设置数据保存到哪个Sheet, 如果不指定,就是保存到默认的Sheet1
}
