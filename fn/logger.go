package fn

import (
	"log"
)

// Logger : write info into Console OR File
func Logger(format string, v ...interface{}) {

	typ := mFnType[caller(false)]

	switch {
	case log2f && log2c:
		// FILE
		v1 := append([]interface{}{typ}, v...)
		item := fSf("\t%s \t\""+format+"\"\n", v1...)
		log.Printf("%s", item)
		// CONSOLE
		v2 := append([]interface{}{green(typ)}, v...)
		item = fSf("\t%s \t\""+format+"\"\n", v2...)
		fPt(tmstr() + item)

	case log2f && !log2c:
		// FILE
		v1 := append([]interface{}{typ}, v...)
		item := fSf("\t%s \t\""+format+"\"\n", v1...)
		log.Printf("%s", item)

	case !log2f && log2c:
		// CONSOLE
		v1 := append([]interface{}{green(typ)}, v...)
		item := fSf("\t%s \t\""+format+"\"\n", v1...)
		log.Printf("%s", item)
	}
}

// LoggerWhen : write info into Console OR File
func LoggerWhen(condition bool, format string, v ...interface{}) {
	if condition {
		Logger(format, v...)
	}
}
