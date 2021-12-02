package xiu

import (
	"testing"
	"xiuexcel/config"
	"xiuexcel/loglib"

	"github.com/sirupsen/logrus"
)

func init() {
	config.Init("test")
	loglib.Init()
}
func TestRead(t *testing.T) {
	log := loglib.GetLog()

	c := config.GetConfig()
	sourceFile := c.GetString("excel.sourceFile")
	sourceSheet := c.GetString("excel.sourceSheet")
	destinationDirectory := c.GetString("excel.destinationDirectory")

	xiuExcel := NewXiuExcel(sourceFile, sourceSheet, destinationDirectory)
	cols, err := xiuExcel.Read()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("read excel file failed.")

	}

	colIndex := 0
	for cols.Next() {
		col, err := cols.Rows()
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("error.")
		}

		for index, rowCell := range col {
			log.WithFields(logrus.Fields{
				"index":   index,
				"rowCell": rowCell,
			}).Info("Row Cell Content.")
		}

		colIndex++
	}

}
