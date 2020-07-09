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
					v = append([]interface{}{mFnType[caller(false)]}, v...)
					warnItem := fSf("\t%s \t\""+format+"\"%s\n", append(v, tc)...)
					log.Printf(warnItem)
					return tc, fEf("%v", tmstr()+rmTailFromLast(warnItem, tc))
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

// WarnPOnErr : write error into Console OR File
func WarnPOnErr(format string, v ...interface{}) (string, error) {
	return warnOnErr(5, format, v...)
}

// WarnPOnErrWhen : write error into Console OR File
func WarnPOnErrWhen(condition bool, format string, v ...interface{}) (string, error) {
	if condition {
		return warnOnErr(5, format, v...)
	}
	return "", nil
}