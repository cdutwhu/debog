package fn

import (
	"log"
)

func debug(lvl int, format string, v ...interface{}) {
	tc := trackCaller(lvl)
	typ := mFnType[caller(false)]

	switch {
	case log2F && log2C:
		// FILE
		v1 := append([]interface{}{typ}, v...)
		item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
		log.Printf("%s", item)
		// CONSOLE
		v2 := append([]interface{}{B(typ)}, v...)
		item = fSf("\t%s \t\""+format+"\"%s", append(v2, tc)...)
		fPt(tmstr() + item)

	case log2F && !log2C:
		// FILE
		v1 := append([]interface{}{typ}, v...)
		item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
		log.Printf("%s", item)

	case !log2F && log2C:
		// CONSOLE
		v1 := append([]interface{}{B(typ)}, v...)
		item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
		log.Printf("%s", item)
	}
}

// Debug : write info into Console OR File
func Debug(format string, v ...interface{}) {
	debug(2, format, v...)
}

// DebugWhen : write info into Console OR File
func DebugWhen(condition bool, format string, v ...interface{}) {
	if condition {
		debug(2, format, v...)
	}
}

// DebugP1 :
func DebugP1(format string, v ...interface{}) {
	debug(3, format, v...)
}

// DebugP1When : write info into Console OR File
func DebugP1When(condition bool, format string, v ...interface{}) {
	if condition {
		debug(3, format, v...)
	}
}
