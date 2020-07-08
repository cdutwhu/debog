package fn

import (
	"testing"
)

func TestLogger(t *testing.T) {
	SetLog(logfile4test)
	str := Logger("%v", caller(false))
	fPln(str)
}

func TestLoggerWhen(t *testing.T) {
	SetLog(logfile4test)
	str := LoggerWhen(true, "%v", caller(false))
	fPln(str)
}
