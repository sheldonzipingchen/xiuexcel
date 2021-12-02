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
	_, err := xiuExcel.Read()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("read excel file failed.")

	}

}

func TestWrite(t *testing.T) {
	c := config.GetConfig()
	sourceFile := c.GetString("excel.sourceFile")
	sourceSheet := c.GetString("excel.sourceSheet")
	destinationDirectory := c.GetString("excel.destinationDirectory")

	xiuExcel := NewXiuExcel(sourceFile, sourceSheet, destinationDirectory)
	cols, _ := xiuExcel.Read()
	xiuExcel.Write(cols)
}
