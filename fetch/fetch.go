package fetch

import (
	"io/ioutil"
	"regexp"
	"time"

	"github.com/cdutwhu/debog/base"
)

// GetLog : logType [INFO, DEBUG, WARN, FAIL]; tmBackwards second unit
func GetLog(logFile, logType string, tmBackwards int, desc bool) ([]string, error) {

	if !base.Exist(logType, "INFO", "DEBUG", "WARN", "FAIL") {
		return nil, fEf("logType is not supported")
	}

	// LIKE `ABC@AEST+ 2.0.log` `DEF@AEST-10.0.log`
	r := regexp.MustCompile(`.+@.+[\+-][ 0-9]{2}\.[0-9]\.log$`)
	if !r.MatchString(logFile) {
		return nil, fEf("logFile is not acceptable format")
	}

	bytes, err := ioutil.ReadFile(logFile)
	// FailOnErr("%v", err)
	if err != nil {
		return nil, err
	}

	logFileHead := logFile[:len(logFile)-4]     // remove ".log"
	offset := logFileHead[len(logFileHead)-5:]  // pick up offset hour after "@", [+/-]NN.N
	offset = sReplaceAll(offset, " ", "")       // remove BLANK " " in such as "+ 2.0"
	hourOffset, err := scParseFloat(offset, 32) // string to float
	if err != nil {
		return nil, err
	}
	tmOffset := hourOffset * 3600 // hours to seconds

	past := time.Now().UTC().Add(-time.Second * time.Duration(tmBackwards))
	fPln(past)

	past = time.Now().UTC()
	fPln(past)
	fPln(time.Now())

	fPln(" ------------------ ")

	re := regexp.MustCompile(fSf(`^[0-9/: ]{20}\t%s \t`, logType))
	logs := []string{}
	for _, ln := range sSplit(string(bytes), "\n") {
		if re.MatchString(ln) {
			tm, err := time.Parse(tmFmt, ln[:19])
			if err != nil {
				return nil, err
			}
			// FailOnErr("%v", err)
			if tm.After(past) {
				tm = tm.Add(time.Second * time.Duration(tmOffset))
				ln = tm.Format(tmFmt) + ln[19:]
				logs = append(logs, ln)
			}
		}
	}

	if desc {
		for l, r := 0, len(logs)-1; l < r; l, r = l+1, r-1 {
			logs[l], logs[r] = logs[r], logs[l]
		}
	}

	return logs, nil
}

// Log2CSV :
func Log2CSV(logFile, logType string, tmBackwards int, desc bool) (string, error) {
	logs, err := GetLog(logFile, logType, tmBackwards, desc)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", nil
	}
	content := "Time,Type,Desc\n"
	content += sReplaceAll(sJoin(logs, "\n"), " \t", ",")
	file := base.RmTailFromLast(logFile, ".") + "-" + logType + ".csv"
	base.MustWriteFile(file, []byte(content))
	return file, nil
}

// Log2File :
func Log2File(logFile, logType string, tmBackwards int, desc bool) (string, error) {
	logs, err := GetLog(logFile, logType, tmBackwards, desc)
	if err != nil {
		return "", err
	}
	if len(logs) == 0 {
		return "", nil
	}
	content := sJoin(logs, "\n")
	file := base.RmTailFromLast(logFile, ".") + "-" + logType + "." + base.RmHeadToLast(logFile, ".")
	base.MustWriteFile(file, []byte(content))
	return file, nil
}
