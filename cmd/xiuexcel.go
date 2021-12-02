package main

import (
	"flag"
	"fmt"
	"os"
	"xiuexcel/config"
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
}

func main() {
	xiuExcelMain()
	os.Exit(exitCode)
}