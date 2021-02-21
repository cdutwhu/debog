package fetch

import (
	"os"
	"regexp"
	"time"
)

// GetLog : logType [INFO, DEBUG, WARN, FAIL]; tmBackwards second unit
func GetLog(logFile, logType string, tmBackwards int, desc bool) []string {

	failP1OnErrWhen(!exist(logType, "ALL", "INFO", "DEBUG", "WARN", "FAIL"), "%v", fEf("logType is not supported"))

	// LIKE `ABC@AEST+ 2.0.log` `DEF@AEST-10.0.log`
	r := regexp.MustCompile(`.+@.+[\+-][ 0-9]{2}\.[0-9]\.log$`)
	failP1OnErrWhen(!r.MatchString(logFile), "%v", fEf("logFile is not acceptable format"))

	bytes, err := os.ReadFile(logFile)
	failP1OnErr("%v", err)

	logFileHead := logFile[:len(logFile)-4]     // remove ".log"
	offset := logFileHead[len(logFileHead)-5:]  // pick up offset hour after "@", [+/-]NN.N
	offset = sReplaceAll(offset, " ", "")       // remove BLANK " " in such as "+ 2.0"
	hourOffset, err := scParseFloat(offset, 32) // string to float
	failOnErr("%v", err)
	tmOffset := hourOffset * 3600 // hours to seconds
	// fPln(tmOffset)

	nowUTC := time.Now().UTC()
	// fPln(nowUTC)
	pastUTC := nowUTC.Add(-time.Second * time.Duration(tmBackwards))
	// fPln(pastUTC)
	// fPln(" ------------------ ")

	re := regexp.MustCompile(fSf(`^[0-9/: ]{20}\t%s \t`, logType))
	if logType == "ALL" {
		re = regexp.MustCompile(`^[0-9/: ]{20}\t`)
	}

	logs := []string{}
	for _, ln := range sSplit(string(bytes), "\n") {
		if re.MatchString(ln) {
			tm, err := time.Parse(tmFmt, ln[:19])
			failOnErr("%v %v", tm, err)
			tmUTC := tm.Add(-time.Second * time.Duration(tmOffset)) // log items -> UTC time
			// fPln(tmUTC)
			if tmUTC.After(pastUTC) {
				logs = append(logs, ln)
			}
		}
	}

	if desc {
		for l, r := 0, len(logs)-1; l < r; l, r = l+1, r-1 {
			logs[l], logs[r] = logs[r], logs[l]
		}
	}

	return logs
}

// Log2CSV :
func Log2CSV(logFile, logType string, tmBackwards int, desc bool) string {
	logs := GetLog(logFile, logType, tmBackwards, desc)
	if len(logs) == 0 {
		return ""
	}
	content := "Time,Type,Desc\n"
	content += sReplaceAll(sJoin(logs, "\n"), " \t", ",")
	file := rmTailFromLast(logFile, "@") + "-" + logType + "@" + rmHeadToLast(logFile, "@")
	file = file[:len(file)-4] + ".csv"
	mustWriteFile(file, []byte(content))
	return file
}

// Log2File :
func Log2File(logFile, logType string, tmBackwards int, desc bool) string {
	logs := GetLog(logFile, logType, tmBackwards, desc)
	if len(logs) == 0 {
		return ""
	}
	content := sJoin(logs, "\n")
	file := rmTailFromLast(logFile, "@") + "-" + logType + "@" + rmHeadToLast(logFile, "@")
	mustWriteFile(file, []byte(content))
	return file
}
