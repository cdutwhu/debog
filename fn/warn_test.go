package fn

import (
	"testing"
)

func TestWarnOnErr(t *testing.T) {
	EnableLog2F(true, logfile4test)
	err := WarnOnErr("%v", fEf(caller(false)))
	fPln(err)
}

func TestWarnOnErrWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	err := WarnOnErrWhen(false, "%v", fEf(caller(false)))
	fPln(err)
	fPln(" -------------------- ")
	err = WarnOnErrWhen(true, "%v", fEf(caller(false)))
	fPln(err)
}

func FakeFuncWarnP(i int) {
	err := WarnP1OnErrWhen(i < 0, "%v", fEf("Invalid Param"))
	fPln(err)
}

func TestWarnP1OnErrWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	FakeFuncWarnP(-5)
}
