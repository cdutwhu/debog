package fn

import (
	"testing"
)

func TestDebug(t *testing.T) {
	SetLog(logfile4test)
	str := Debug("%v", caller(false))
	fPln("---" + str)
	ResetLog()
	Debug("%v", caller(false))
}

func TestDebugWhen(t *testing.T) {
	SetLog(logfile4test)
	str := DebugWhen(true, "%v", caller(false))
	fPln("---" + str)
}

func FakeFuncDebugP(i int) {
	where := DebugP1When(i < 0, "Invalid Param")
	fPln("---" + where)
}

func TestDebugPWhen(t *testing.T) {
	SetLog(logfile4test)
	FakeFuncDebugP(-5)
}
