package fn

import (
	"testing"
)

func TestLogger(t *testing.T) {
	EnableLog2F(true, logfile4test)
	Logger("%v", caller(false))	
}

func TestLoggerWhen(t *testing.T) {
	EnableLog2F(true, logfile4test)
	LoggerWhen(true, "%v", caller(false))	
}
