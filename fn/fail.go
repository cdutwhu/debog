package fn

import (
	"log"
)

func failOnErr(lvl int, format string, v ...interface{}) {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					tc := trackCaller(lvl)
					typ := mFnType[caller(false)]

					switch {
					case log2F && log2C:
						// CONSOLE
						v1 := append([]interface{}{R(typ)}, v...)
						fatalInfo := fSf("\t%s \t\""+format+"\"%s\n", append(v1, tc)...)
						fPt(tmstr() + fatalInfo)
						// FILE
						v2 := append([]interface{}{typ}, v...)
						fatalInfo = fSf("\t%s \t\""+format+"\"%s\n", append(v2, tc)...)
						log.Fatalf(fatalInfo)

					case log2F && !log2C:
						// FILE
						v2 := append([]interface{}{typ}, v...)
						fatalInfo := fSf("\t%s \t\""+format+"\"%s\n", append(v2, tc)...)
						log.Fatalf(fatalInfo)

					case !log2F && log2C:
						// CONSOLE
						v2 := append([]interface{}{R(typ)}, v...)
						fatalInfo := fSf("\t%s \t\""+format+"\"%s\n", append(v2, tc)...)
						log.Fatalf(fatalInfo)
					}
				}
			}
		}
	}
}

// FailOnErr : error holder use "%v"
func FailOnErr(format string, v ...interface{}) {
	failOnErr(2, format, v...)
}

// FailOnErrWhen :
func FailOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		failOnErr(2, format, v...)
	}
}

// FailP1OnErr : error holder use "%v"
func FailP1OnErr(format string, v ...interface{}) {
	failOnErr(3, format, v...)
}

// FailP1OnErrWhen :
func FailP1OnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		failOnErr(3, format, v...)
	}
}

// FailPnOnErr : error holder use "%v"
func FailPnOnErr(n int, format string, v ...interface{}) {
	failOnErr(2+n, format, v...)
}

// FailPnOnErrWhen :
func FailPnOnErrWhen(condition bool, n int, format string, v ...interface{}) {
	if condition {
		failOnErr(2+n, format, v...)
	}
}
