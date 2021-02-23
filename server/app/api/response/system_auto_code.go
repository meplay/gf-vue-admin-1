package response

type Dbs struct {
	Database string `json:"database"`
}

type Tables struct {
	TableName string `json:"tableName"`
}

type Columns struct {
	DataType      string `orm:"data_type" json:"dataType"`
	ColumnName    string `orm:"column_name" json:"columnName"`
	DataTypeLong  string `orm:"data_type_long" json:"dataTypeLong"`
	ColumnComment string `orm:"column_comment" json:"columnComment"`
}
