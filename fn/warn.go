package fn

import (
	"log"
)

func warnOnErr(lvl int, format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					tc := trackCaller(lvl)
					typ := mFnType[caller(false)]

					switch {
					case log2f && log2c:
						// FILE
						v1 := append([]interface{}{typ}, v...)
						item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
						log.Printf("%s", item)
						// CONSOLE
						v2 := append([]interface{}{yellow(typ)}, v...)
						item = fSf("\t%s \t\""+format+"\"%s", append(v2, tc)...)
						fPt(tmstr() + item)

					case log2f && !log2c:
						// FILE
						v1 := append([]interface{}{typ}, v...)
						item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
						log.Printf("%s", item)

					case !log2f && log2c:
						// CONSOLE
						v1 := append([]interface{}{yellow(typ)}, v...)
						item := fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
						log.Printf("%s", item)
					}

					return p.(error)
				}
			}
		}
	}
	return nil
}

// WarnOnErr : write error into Console OR File
func WarnOnErr(format string, v ...interface{}) error {
	return warnOnErr(4, format, v...)
}

// WarnOnErrWhen : write error into Console OR File
func WarnOnErrWhen(condition bool, format string, v ...interface{}) error {
	if condition {
		return warnOnErr(4, format, v...)
	}
	return nil
}

// WarnP1OnErr : write error into Console OR File
func WarnP1OnErr(format string, v ...interface{}) error {
	return warnOnErr(5, format, v...)
}

// WarnP1OnErrWhen : write error into Console OR File
func WarnP1OnErrWhen(condition bool, format string, v ...interface{}) error {
	if condition {
		return warnOnErr(5, format, v...)
	}
	return nil
}
