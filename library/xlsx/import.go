package xlsx

import (
	_errors "github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

type Import interface {
	Rows2Database(row [][]string) error
}

type _import struct {
	rows      [][]string
	filePath  string
	sheetName string
	header    *multipart.FileHeader
}

func NewImport() *_import {
	return &_import{}
}

func (i *_import) SetSheetName(name string) *_import {
	if name == "" {
		i.sheetName = "Sheet1"
	}
	i.sheetName = name
	return i
}

func (i *_import) SetHeader(header *multipart.FileHeader) *_import {
	i.header = header
	return i
}

func (i *_import) Import(info Import) error {
	file, err := i.header.Open()
	if err != nil {
		return _errors.Wrap(err, "文件流打开失败!")
	}
	_excel, err := excelize.OpenReader(file)
	if err != nil {
		return _errors.Wrap(err, "excelize 库打开文件流失败!")
	}
	rows, err := _excel.GetRows(i.sheetName)
	if err != nil {
		return _errors.Wrap(err, "获取 表格数据失败!")
	}
	if err = info.Rows2Database(rows); err != nil {
		return _errors.Wrap(err, "数据入库失败!")
	}
	return nil
}
