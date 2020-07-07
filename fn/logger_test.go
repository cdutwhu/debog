package fn

import (
	"testing"

	"github.com/cdutwhu/debog/base"
)

func TestLogger(t *testing.T) {
	SetLog(logfile4test)
	str := Logger("%v", base.Caller(false))
	fPln(str)
}

func TestLoggerWhen(t *testing.T) {
	SetLog(logfile4test)
	str := LoggerWhen(true, "%v", base.Caller(false))
	fPln(str)
}
