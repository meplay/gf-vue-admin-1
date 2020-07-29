package request

type DBReq struct {
	Database string `json:"database"`
}

type TableReq struct {
	TableName string `json:"tableName"`
}

type ColumnReq struct {
	ColumnName    string `json:"columeName"`
	DataType      string `json:"dataType"`
	DataTypeLong  string `json:"dataTypeLong"`
	ColumnComment string `json:"columeComment"`
}
