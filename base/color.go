package base

import (
	"regexp"
)

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	TEAL
	WHITE
	CLRCOUNT
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fSf(colorString, fSt(args...))
	}
	return sprint
}

// ClrGrp :
var ClrGrp = []func(...interface{}) string{
	color("\033[1;30m%s\033[0m"),
	color("\033[1;31m%s\033[0m"),
	color("\033[1;32m%s\033[0m"),
	color("\033[1;33m%s\033[0m"),
	color("\033[1;34m%s\033[0m"),
	color("\033[1;35m%s\033[0m"),
	color("\033[1;36m%s\033[0m"),
	color("\033[1;37m%s\033[0m"),
}

// DeColor :
func DeColor(str string) string {
	r := regexp.MustCompile("\\033\\[1;3[0-7]m")
	str = sReplaceAll(str, "\033[0m", "")
	return string(r.ReplaceAll([]byte(str), []byte("")))
}

// R :
func R(s string) string {
	return ClrGrp[RED](s)
}

// G :
func G(s string) string {
	return ClrGrp[GREEN](s)
}

// B :
func B(s string) string {
	return ClrGrp[BLUE](s)
}

// Y :
func Y(s string) string {
	return ClrGrp[YELLOW](s)
}

// W :
func W(s string) string {
	return ClrGrp[WHITE](s)
}
