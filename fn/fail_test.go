package fn

import (
	"testing"
)

func TestFailOnErr(t *testing.T) {
	SetLog(logfile4test)
	FailOnErr("%v", fEf(caller(false)))
}

func TestFailOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	FailOnErrWhen(false, "%v", fEf(caller(false)))
	fPln(" -------------------- ")
	FailOnErrWhen(true, "%v", fEf(caller(false)))
}

func FakeFuncFailP(i int) {
	FailPOnErrWhen(i < 0, "%v", fEf("Invalid Param"))
}

func TestFailPOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	FakeFuncFailP(-5)
}
