package fn

import (
	"log"
)

func warnOnErr(lvl int, format string, v ...interface{}) (string, error) {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					tc := trackCaller(lvl)
					typ := mFnType[caller(false)]
					if !log2file {
						typ = yellow(typ)
					}
					v = append([]interface{}{typ}, v...)
					warnItem := fSf("\t%s \t\""+format+"\"%s\n", append(v, tc)...)
					log.Printf(warnItem)
					return tc, fEf("%v", tmstr()+rmTailFromLast(decolor(warnItem), tc))
				}
			}
		}
	}
	return "", nil
}

// WarnOnErr : write error into Console OR File
func WarnOnErr(format string, v ...interface{}) (string, error) {
	return warnOnErr(4, format, v...)
}

// WarnOnErrWhen : write error into Console OR File
func WarnOnErrWhen(condition bool, format string, v ...interface{}) (string, error) {
	if condition {
		return warnOnErr(4, format, v...)
	}
	return "", nil
}

// WarnP1OnErr : write error into Console OR File
func WarnP1OnErr(format string, v ...interface{}) (string, error) {
	return warnOnErr(5, format, v...)
}

// WarnP1OnErrWhen : write error into Console OR File
func WarnP1OnErrWhen(condition bool, format string, v ...interface{}) (string, error) {
	if condition {
		return warnOnErr(5, format, v...)
	}
	return "", nil
}
