package fn

import (
	"testing"
)

func TestDebug(t *testing.T) {
	EnableLog2F(true, logfile4test)
	Debug("%v", "*** Debug ***")
	EnableLog2F(false, "")
	Debug("%v", "*** Debug ***")
}

func TestDebugWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	DebugWhen(true, "%v", caller(false))
}

func FakeFuncDebugP(i int) {
	DebugP1When(i < 0, "Invalid Param")
}

func TestDebugPWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	FakeFuncDebugP(-5)
}
