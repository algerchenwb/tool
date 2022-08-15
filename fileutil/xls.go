package fileutil

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shakinm/xlsReader/xls"
	"github.com/shakinm/xlsReader/xls/structure"
	"github.com/tealeg/xlsx"
)

func Example(fromXlsPath string, toXlsxPath string) {
	xlsFile, deleteFun, err := Xls2Xlsx(fromXlsPath, toXlsxPath)
	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}
	data, err := xlsFile.ToSlice()
	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}
	for i := range data {
		for j := range data[i] {
			for k, v := range data[i][j] {
				fmt.Println(k, v)
			}
		}
	}
	time.Sleep(10 * time.Second)
	err = deleteFun()
	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}

}
func Xls2Xlsx(fromXlsPath string, toXlsxPath string) (*xlsx.File, func() error, error) { // 打開xls文件
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v", err)
		}
	}()
	fromXlsSheet, err := openXlsFile(fromXlsPath)
	if err != nil {
		return nil, nil, err
	}

	// 打開xlsx文件
	toXlsxFile, err := openXlsxFile(toXlsxPath)
	if err != nil {
		return nil, nil, err
	}
	toXlsxSheet, err := toXlsxFile.AddSheet("0")
	if err != nil {
		return nil, nil, err
	}

	// 数据拷贝
	err = dataCopy(fromXlsSheet, toXlsxSheet)
	if err != nil {
		return nil, nil, err
	}

	// 保存文件
	err = toXlsxFile.Save(toXlsxPath)
	if err != nil {
		return nil, nil, err
	}

	deleteFun := func() error {
		err := os.Remove(toXlsxPath)
		return err
	}
	return toXlsxFile, deleteFun, nil
}

func openXlsFile(fromXlsPath string) (*xls.Sheet, error) {
	// xls文件存在
	if !IsFile(fromXlsPath) {
		return nil, fmt.Errorf("file %s not found", fromXlsPath)
	}

	//是xls文件
	if !strings.HasSuffix(fromXlsPath, ".xls") {
		return nil, fmt.Errorf("file %s not a xls file", fromXlsPath)
	}

	// open xls file
	fromXlsFile, err := xls.OpenFile(fromXlsPath)
	if err != nil {
		return nil, err
	}
	fromXlsSheet, err := fromXlsFile.GetSheet(0)
	return fromXlsSheet, nil
}

func dataCopy(fromXlsSheet *xls.Sheet, toXlsxSheet *xlsx.Sheet) error {
	for i := 0; i < int(fromXlsSheet.GetNumberRows()); i++ {
		fromXlsRow, err := fromXlsSheet.GetRow(i)
		if err != nil {
			return err
		}
		rowCols := fromXlsRow.GetCols()
		insertRowFrom(toXlsxSheet, rowCols)
	}
	return nil
}

func insertRowFrom(toXlsxSheet *xlsx.Sheet, rowCols []structure.CellData) {

	row := toXlsxSheet.AddRow()
	for _, col := range rowCols {
		cell := row.AddCell()
		cell.Value = col.GetString()
	}
}

func openXlsxFile(toXlsxPath string) (*xlsx.File, error) {
	err := os.Remove(toXlsxPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	xlsxFile := xlsx.NewFile()

	return xlsxFile, nil
}
