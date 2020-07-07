package fn

import (
	"testing"

	"github.com/cdutwhu/debog/base"
)

func TestWarnOnErr(t *testing.T) {
	SetLog(logfile4test)
	where, err := WarnOnErr("%v", fEf(base.Caller(false)))
	fPln(where, err)
}

func TestWarnOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	where, err := WarnOnErrWhen(false, "%v", fEf(base.Caller(false)))
	fPln(err, where)
	fPln(" -------------------- ")
	where, err = WarnOnErrWhen(true, "%v", fEf(base.Caller(false)))
	fPln(err, where)
}

func FakeFuncWarnP(i int) {
	where, err := WarnPOnErrWhen(i < 0, "%v", fEf("Invalid Param"))
	fPln(err, where)
}

func TestWarnPOnErrWhen(t *testing.T) {
	SetLog(logfile4test)
	FakeFuncWarnP(-5)
}
