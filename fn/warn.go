package fn

import (
	"log"
)

const (
	toF = iota
	toC
)

func warnOnErr(lvl int, format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					tc := trackCaller(lvl)
					typ := mFnType[caller(false)]
					item := ""

					switch {
					case log2F && log2C:
						for i := toF; i <= toC; i++ {
							if i == toC {
								typ = yellow(typ)
							}
							v1 := append([]interface{}{typ}, v...)
							if warnWithDetail {
								item = fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
							} else {
								item = fSf("\t%s \t\""+format+"\"\n", v1...)
							}
							if i == toF {
								log.Printf("%s", item)
							} else if i == toC {
								fPt(tmstr() + item)
							}
						}
						return p.(error)

					case log2F && !log2C:
					case !log2F && log2C:
						typ = yellow(typ)
					}

					v1 := append([]interface{}{typ}, v...)
					if warnWithDetail {
						item = fSf("\t%s \t\""+format+"\"%s", append(v1, tc)...)
					} else {
						item = fSf("\t%s \t\""+format+"\"\n", v1...)
					}
					log.Printf("%s", item)
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
