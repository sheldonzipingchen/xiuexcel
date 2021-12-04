package main

import (
	"flag"
	"fmt"
	"os"
	"xiuexcel/config"
	"xiuexcel/loglib"
	"xiuexcel/xiu"

	"github.com/sirupsen/logrus"
)

var (
	// main operation modes
	environment = flag.String("e", "development", "application run in which environment")
)

var (
	exitCode = 0
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: xiuexcel [flags]\n")
	flag.PrintDefaults()
}

func xiuExcelMain() {
	flag.Usage = usage
	flag.Parse()

	config.Init(*environment)
	loglib.Init()

	log := loglib.GetLog()
	c := config.GetConfig()

	sourceFile := c.GetString("excel.sourceFile")
	sourceSheet := c.GetString("excel.sourceSheet")
	destinationDirectory := c.GetString("excel.destinationDirectory")

	xiuExcel := xiu.NewXiuExcel(sourceFile, sourceSheet, destinationDirectory)
	cols, err := xiuExcel.Read()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("read excel file failed.")

	}

	xiuExcel.Write(cols)
}

func main() {
	xiuExcelMain()
	os.Exit(exitCode)
}
