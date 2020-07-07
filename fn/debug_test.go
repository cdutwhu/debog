package fn

import (
	"testing"

	"github.com/cdutwhu/debog/base"
)

func TestDebug(t *testing.T) {
	SetLog(logfile4test)
	str := Debug("%v", base.Caller(false))
	fPln("---" + str)
	ResetLog()
	Debug("%v", base.Caller(false))
}

func TestDebugWhen(t *testing.T) {
	SetLog(logfile4test)
	str := DebugWhen(true, "%v", base.Caller(false))
	fPln("---" + str)
}

func FakeFuncDebugP(i int) {
	where := DebugPWhen(i < 0, "Invalid Param")
	fPln("---" + where)
}

func TestDebugPWhen(t *testing.T) {
	SetLog(logfile4test)
	FakeFuncDebugP(-5)
}
