package xiu

import (
	"fmt"
	"xiuexcel/loglib"

	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type XiuExcel struct {
	sourceFile           string // 源文件路径（绝对路径）
	sourceSheet          string // 源文件 Sheet 名称
	destinationDirectory string // 生成文件目录名称
}

func NewXiuExcel(sourceFile string, sourceSheet string, destinationDirectory string) *XiuExcel {
	return &XiuExcel{
		sourceFile:           sourceFile,
		sourceSheet:          sourceSheet,
		destinationDirectory: destinationDirectory,
	}
}

// 读取文件内容
func (x *XiuExcel) Read() (*excelize.Cols, error) {
	log := loglib.GetLog()

	f, err := excelize.OpenFile(x.sourceFile)

	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Open Excel File Failed.")

		return nil, err
	}

	cols, err := f.Cols(x.sourceSheet)

	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Read Excel File Failed.")

		return nil, err
	}

	return cols, nil
}

// Write 输出结果文件
func (x *XiuExcel) Write(cols *excelize.Cols) {
	log := loglib.GetLog()

	colIndex := 0

	firstCol := []string{}
	secondCol := []string{}

	for cols.Next() {

		rows, err := cols.Rows()
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("error.")
		}

		result := []string{}
		result = append(result, rows...)

		if colIndex == 0 {
			firstCol = result
		}

		if colIndex == 1 {
			secondCol = result
		}

		if colIndex > 2 {
			fileName := fmt.Sprintf("%s%s.xlsx", x.destinationDirectory, result[0])
			f := excelize.NewFile()

			for index, row := range firstCol {
				colName := fmt.Sprintf("A%d", index+1)
				f.SetCellValue("Sheet1", colName, row)
			}

			for index, row := range secondCol {
				colName := fmt.Sprintf("B%d", index+1)
				f.SetCellValue("Sheet1", colName, row)
			}

			for index, row := range result {
				colName := fmt.Sprintf("C%d", index+1)
				f.SetCellValue("Sheet1", colName, row)
			}

			if err := f.SaveAs(fileName); err != nil {
				log.WithFields(logrus.Fields{
					"fileName": fileName,
					"error":    err,
				}).Error("Save Excel File Error.")
			}
		}

		colIndex++
	}
}
