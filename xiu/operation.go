package xiu

import (
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
func (x *XiuExcel) Write(*excelize.Cols) {

}
