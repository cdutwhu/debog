package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
)

func main() {

	// for i := 0; i < base.CLRCOUNT; i++ {
	// 	fmt.Println(base.ClrGrp[i]("color"), "normal")
	// }
	// fmt.Println(base.R("Red"))
	// fmt.Println(base.G("Green"))
	// fmt.Println(base.B("Blue"))
	// fmt.Println(base.Y("Yellow"))
	// fmt.Println(base.W("White"))

	// fmt.Println(" ------------------------------------- ")

	fn.EnableLog2F(true, "./a.log")
	fn.EnableLog2C(true)
	fn.EnableWarnDetail(false)

	fn.Logger("logger test")
	fn.Debug("debug test")
	fn.WarnOnErr("%v", fmt.Errorf("warn test"))
	fn.FailOnErr("%v", fmt.Errorf("fail test"))
}
