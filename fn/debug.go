package fn

import (
	"log"
)

func debug(lvl int, format string, v ...interface{}) string {
	tc := trackCaller(lvl)
	v = append([]interface{}{mFnType[caller(false)]}, v...)
	logItem := fSf("\t%s \t\""+format+"\"%s\n", append(v, tc)...)
	log.Printf("%s", logItem)
	return tmstr() + logItem
}

// Debug : write info into Console OR File
func Debug(format string, v ...interface{}) string {
	return debug(4, format, v...)
}

// DebugWhen : write info into Console OR File
func DebugWhen(condition bool, format string, v ...interface{}) string {
	if condition {
		return debug(4, format, v...)
	}
	return ""
}

// DebugP :
func DebugP(format string, v ...interface{}) string {
	return debug(5, format, v...)
}

// DebugPWhen : write info into Console OR File
func DebugPWhen(condition bool, format string, v ...interface{}) string {
	if condition {
		return debug(5, format, v...)
	}
	return ""
}
