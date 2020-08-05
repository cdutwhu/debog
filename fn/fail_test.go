package fn

import (
	"testing"
)

func TestFailOnErr(t *testing.T) {
	EnableLog2F(true, logfile4test)
	FailOnErr("%v", fEf(caller(false)))
}

func TestFailOnErrWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	FailOnErrWhen(false, "%v", fEf(caller(false)))
	fPln(" -------------------- ")
	FailOnErrWhen(true, "%v", fEf(caller(false)))
}

func FakeFuncFailP(i int) {
	FailP1OnErrWhen(i < 0, "%v", fEf("Invalid Param"))
}

func TestFailP1OnErrWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	FakeFuncFailP(-5)
}
