package system

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type AutoCodeInjections []AutoCodeInjection

type AutoCodeInjection struct {
	Filepath       string `json:"filepath" gorm:"column:filepath;comment:注入的文件绝对路径"`
	CodeData       string `json:"codeData" gorm:"column:code_data;comment:注入代码"`
	StructName     string `json:"structName" gorm:"column:struct_name;comment:注入的结构体名"`
	FunctionName   string `json:"functionName" gorm:"column:function_name;comment:注入的函数名"`
	CodeDataFormat string `json:"codeDataFormat" gorm:"column:code_data_format;comment:注入代码带格式化"`
}

// Scan 扫描
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeInjections) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		if err := json.Unmarshal(v, a); err != nil {
			return err
		}
	case string:
		if err := json.Unmarshal([]byte(v), a); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal AutoCodeInjection value:", value))
	}
	return nil
}

// Value 值
// Author [SliverHorn](https://github.com/SliverHorn)
func (a AutoCodeInjections) Value() (driver.Value, error) {
	bytes, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}
	return driver.Value(bytes), err
}
