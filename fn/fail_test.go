package fn

import (
	"testing"

	"github.com/cdutwhu/debog/base"
)

func TestFailOnErr(t *testing.T) {
	SetLog(logfile4test)
	FailOnErr("%v", fEf(base.Caller(false)))
}

func TestFailOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	FailOnErrWhen(false, "%v", fEf(base.Caller(false)))
	fPln(" -------------------- ")
	FailOnErrWhen(true, "%v", fEf(base.Caller(false)))
}

func FakeFuncFailP(i int) {
	FailPOnErrWhen(i < 0, "%v", fEf("Invalid Param"))
}

func TestFailPOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	FakeFuncFailP(-5)
}
