package fn

import (
	"log"
)

// Logger : write info into Console OR File
func Logger(format string, v ...interface{}) string {
	typ := mFnType[caller(false)]
	if !log2file {
		typ = green(typ)
	}
	v = append([]interface{}{typ}, v...)
	logItem := fSf("\t%s \t\""+format+"\"\n\n", v...)
	log.Printf("%s", logItem)
	return tmstr() + decolor(logItem)
}

// LoggerWhen : write info into Console OR File
func LoggerWhen(condition bool, format string, v ...interface{}) string {
	if condition {
		return Logger(format, v...)
	}
	return ""
}
