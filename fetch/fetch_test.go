package fetch

import (
	"testing"

	"github.com/cdutwhu/debog/fn"
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

	logs, err := GetLog("../a/b@AEST+11.0.log", "WARN", 10, true)
	fn.FailOnErr("%v", err)
	for _, ln := range logs {
		fPln(ln)
	}
	// FetchLog2File("./mylog@AEST+10.log", "FAIL", 10000, true)
	// FetchLog2CSV("./mylog@AEST+10.log", "FAIL", 10000, true)
}
