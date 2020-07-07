package base

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			log.Fatalf("Could NOT Create File to Write: %v @ %s", err, Caller(true))
		}
	}
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		log.Fatalf("Could NOT Write File: %v @ %s", err, Caller(true))
	}
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		MustWriteFile(filename, []byte(""))
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could NOT Open File to Write: %v @ %s", err, Caller(true))
	}
	defer file.Close()

	if newline {
		data = append(append([]byte{}, '\n'), data...)
	}
	if _, err = file.Write(data); err != nil {
		log.Fatalf("Could NOT Write File: %v @ %s", err, Caller(true))
	}
}

// Exist :
func Exist(e interface{}, set ...interface{}) bool {
	v := reflect.ValueOf(set)
	l := v.Len()
	for i := 0; i < l; i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

// RmTailFromLast :
func RmTailFromLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmHeadToLast :
func RmHeadToLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[i+len(mark):]
	}
	return s
}

// ------------------------------------------------------------------------------ //

// Caller : get the caller name or path which called this `Caller`
func Caller(fullpath bool) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	if fullpath {
		return frame.Function
	}
	return RmHeadToLast(frame.Function, ".")
}

// ParentCaller : get the caller's caller name or path which called this `ParentCaller`
func ParentCaller(fullpath bool) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	if fullpath {
		return frame.Function
	}
	return RmHeadToLast(frame.Function, ".")
}

// FuncTrack : full path of func name
func FuncTrack(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// TrackCaller :
func TrackCaller(lvl int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(lvl, pc) // lvl: 3 is for util-FailLog. 2 is for "TrackCaller" caller
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fSf("\n%s:%d\n%s\n", frame.File, frame.Line, frame.Function)
}
