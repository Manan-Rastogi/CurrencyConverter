package helpers

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func ErrorLogF(err string) {
	log.Fatal(color.RedString(err))
}

func ErrorLog(err string) {
	log.Println(color.YellowString(err))
}

func InfoLog(info string) {
	log.Println(color.GreenString(info))
}

func ValueLog(varr string, val interface{}) {
	varr = color.MagentaString(varr + ": ")
	val = color.CyanString(fmt.Sprint(val))

	log.Printf("%v %v", varr, val)
}