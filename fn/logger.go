package fn

import (
	"log"
)

// Logger : write info into Console OR File
func Logger(format string, v ...interface{}) string {
	v = append([]interface{}{mFnType[caller(false)]}, v...)
	logItem := fSf("\t%s \t\""+format+"\"\n\n", v...)
	log.Printf("%s", logItem)
	return tmstr() + logItem
}

// LoggerWhen : write info into Console OR File
func LoggerWhen(condition bool, format string, v ...interface{}) string {
	if condition {
		return Logger(format, v...)
	}
	return ""
}
