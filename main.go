package main

import (
	"fmt"
	"io/ioutil"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
)

func main() {
	for i := 0; i < base.CLRCOUNT; i++ {
		fmt.Println(base.ClrGrp[i]("color"), "normal")
	}

	fmt.Println(" -------------------------------- ")

	fmt.Println(base.R("Red"))
	fmt.Println(base.G("Green"))
	fmt.Println(base.B("Blue"))
	fmt.Println(base.Y("Yellow"))
	fmt.Println(base.W("White"))

	fmt.Println(" -------------------------------- ")

	fn.SetLog("./a.log")

	debugInfo := fn.Debug("debug test")
	// debugInfo = base.DeColor(debugInfo)
	fmt.Println(debugInfo)
	ioutil.WriteFile("./a.log", []byte(debugInfo), 0666)

	fn.WarnOnErr("%v", fmt.Errorf("warn test"))
	fn.Logger("logger test")
	fn.FailOnErr("%v", fmt.Errorf("fail test"))
}
