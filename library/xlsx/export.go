package xlsx

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

// maxCharCount 最多26个字符A-Z
const maxCharCount = 26

type export struct {
	a1         []string
	rows       [][]interface{}
	a1Length   int
	rowsLength int
	sheetName  string
}

func NewExport() *export {
	return &export{}
}

func (e *export) SetSheetName(name string) *export {
	e.sheetName = name
	return e
}

func (e *export) SetA1(a1 []string) *export {
	e.a1 = a1
	e.a1Length = len(e.a1)
	return e
}

func (e *export) SetRows(rows [][]interface{}) *export {
	e.rows = rows
	e.rowsLength = len(e.rows)
	return e
}

// Excel 导出Excel文件
// headers 列名切片， 表头
// rows 数据切片，是一个二维数组
func (e *export) Excel() (*excelize.File, error) {
	f := excelize.NewFile()
	sheetIndex := f.NewSheet(e.sheetName)
	maxColumnRowNameLen := 1 + len(strconv.Itoa(e.rowsLength))
	if e.a1Length > maxCharCount {
		maxColumnRowNameLen++
	} else if e.a1Length > maxCharCount*maxCharCount {
		maxColumnRowNameLen += 2
	}
	columnNames := make([][]byte, 0, e.a1Length)
	for i, header := range e.a1 {
		columnName := e.getColumnName(i, maxColumnRowNameLen)
		columnNames = append(columnNames, columnName)
		// 初始化excel表头，这里的index从1开始要注意
		curColumnName := e.getColumnRowName(columnName, 1)
		err := f.SetCellValue(e.sheetName, curColumnName, header)
		if err != nil {
			return nil, err
		}
	}
	for i := 0; i < e.rowsLength; i++ {
		length := len(e.rows[i])
		for j := 0; j < length; j++ { // 从第二行开始
			err := f.SetCellValue(e.sheetName, e.getColumnRowName(columnNames[j], j+2), e.rows[i][j])
			if err != nil {
				return nil, err
			}
		}
	}
	f.SetActiveSheet(sheetIndex)
	return f, nil
}

// getColumnName 生成列名
// Excel的列名规则是从A-Z往后排;超过Z以后用两个字母表示，比如AA,AB,AC;两个字母不够以后用三个字母表示，比如AAA,AAB,AAC
// 这里做数字到列名的映射：0 -> A, 1 -> B, 2 -> C
// maxColumnRowNameLen 表示名称框的最大长度，假设数据是10行，1000列，则最后一个名称框是J1000(如果有表头，则是J1001),是4位
// 这里根据 maxColumnRowNameLen 生成切片，后面生成名称框的时候可以复用这个切片，而无需扩容
func (e *export) getColumnName(column, maxColumnRowNameLen int) []byte {
	const A = 'A'
	if column < maxCharCount {
		// 第一次就分配好切片的容量
		slice := make([]byte, 0, maxColumnRowNameLen)
		return append(slice, byte(A+column))
	} else {
		// 递归生成类似AA,AB,AAA,AAB这种形式的列名
		return append(e.getColumnName(column/maxCharCount-1, maxColumnRowNameLen), byte(A+column%maxCharCount))
	}
}

// getColumnRowName 生成名称框
// Excel的名称框是用A1,A2,B1,B2来表示的，这里需要传入前一步生成的列名切片，然后直接加上行索引来生成名称框，就无需每次分配内存
func (e *export) getColumnRowName(columnName []byte, rowIndex int) (columnRowName string) {
	l := len(columnName)
	columnName = strconv.AppendInt(columnName, int64(rowIndex), 10)
	columnRowName = string(columnName)
	// 将列名恢复回去
	columnName = columnName[:l]
	return
}
