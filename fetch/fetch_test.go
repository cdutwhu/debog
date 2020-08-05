package fetch

import (
	"testing"
)

func TestFetchLog(t *testing.T) {

	// for _, name := range []string{
	// 	"",
	// 	"Local",
	// 	"Asia/Shanghai",
	// 	"America/New_York",
	// 	"Australia/Melbourne",
	// } {
	// 	t, err := TimeIn(time.Now(), name)
	// 	if err == nil {
	// 		fPln(t.Location(), t.Format("15:04"))
	// 	} else {
	// 		fPln(name, "<time unknown>")
	// 	}
	// }

	logs := GetLog("../a@AEST+10.0.log", "DEBUG", 1000, true)
	for _, ln := range logs {
		fPln(ln)
	}

	fPln(" ----- ")

	fPln(Log2File("../a@AEST+10.0.log", "ALL", 10000, true))
	fPln(Log2CSV("../a@AEST+10.0.log", "ALL", 10000, true))
}
