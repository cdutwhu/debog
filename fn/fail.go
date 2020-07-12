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

					// even if log2file, still output a copy to console.
					if log2file {
						typ := red(mFnType[caller(false)])
						v1 := append([]interface{}{typ}, v...)
						fatalInfo := fSf("\t%s \t\""+format+"\"%s\n", append(v1, tc)...)
						fPln(tmstr() + fatalInfo)
					}

					typ := mFnType[caller(false)]
					if !log2file {
						typ = red(typ)
					}
					v2 := append([]interface{}{typ}, v...)
					fatalInfo := fSf("\t%s \t\""+format+"\"%s\n", append(v2, tc)...)
					log.Fatalf(fatalInfo)
				}
			}
		}
	}
}

// FailOnErr : error holder use "%v"
func FailOnErr(format string, v ...interface{}) {
	failOnErr(4, format, v...)
}

// FailOnErrWhen :
func FailOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		failOnErr(4, format, v...)
	}
}

// FailP1OnErr : error holder use "%v"
func FailP1OnErr(format string, v ...interface{}) {
	failOnErr(5, format, v...)
}

// FailP1OnErrWhen :
func FailP1OnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		failOnErr(5, format, v...)
	}
}

// FailPnOnErr : error holder use "%v"
func FailPnOnErr(n int, format string, v ...interface{}) {
	failOnErr(4+n, format, v...)
}

// FailPnOnErrWhen :
func FailPnOnErrWhen(condition bool, n int, format string, v ...interface{}) {
	if condition {
		failOnErr(4+n, format, v...)
	}
}
