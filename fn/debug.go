package fn

import (
	"log"
)

func debug(lvl int, format string, v ...interface{}) string {
	tc := trackCaller(lvl)
	typ := mFnType[caller(false)]
	if !log2file {
		typ = blue(typ)
	}
	v = append([]interface{}{typ}, v...)
	logItem := fSf("\t%s \t\""+format+"\"%s\n", append(v, tc)...)
	log.Printf("%s", logItem)
	return tmstr() + decolor(logItem)
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

// DebugP1 :
func DebugP1(format string, v ...interface{}) string {
	return debug(5, format, v...)
}

// DebugP1When : write info into Console OR File
func DebugP1When(condition bool, format string, v ...interface{}) string {
	if condition {
		return debug(5, format, v...)
	}
	return ""
}
