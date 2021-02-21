package base

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

const (
	// FilePerm :
	FilePerm = 0666
	// DirPerm :
	DirPerm = 0777
)

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {
	dir := filepath.Dir(filename)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, DirPerm); err != nil { // dir must be 0777 to put writes in
			log.Fatalf("Could NOT Create File to Write: %v @ %s", err, Caller(true))
		}
		goto WRITE
	}
	if err != nil {
		log.Fatalf("Cound NOT Get file Status: %v @ %s", err, Caller(true))
	}

WRITE:
	if err := os.WriteFile(filename, data, FilePerm); err != nil {
		log.Fatalf("Could NOT Write File: %v @ %s", err, Caller(true))
	}
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		MustWriteFile(filename, data)
		return
	}
	if err != nil {
		log.Fatalf("Cound NOT Get file Status: %v @ %s", err, Caller(true))
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, FilePerm)
	if err != nil {
		log.Fatalf("Could NOT Open File to Append: %v @ %s", err, Caller(true))
	}
	defer file.Close()

	if newline {
		data = append([]byte{'\n'}, data...)
	}
	if _, err = file.Write(data); err != nil {
		log.Fatalf("Could NOT Append File: %v @ %s", err, Caller(true))
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

// lvl: 0. where `running trackCaller(...) in its caller, such as TrackCaller, Caller`
func trackCaller(lvl int) (string, int, string) {
	lvl += 2
	pc := make([]uintptr, 15)
	n := runtime.Callers(lvl, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Line, frame.Function
}

// TrackCaller :
// lvl: 0, where `running TrackCaller(0)`, who is its caller.
func TrackCaller(lvl int) string {
	file, line, fn := trackCaller(lvl + 1)
	return fSf("\n%s:%d\n%s\n", file, line, fn)
}

// CallerSrc :
func CallerSrc() (dir, src string) {
	file, _, _ := trackCaller(1)
	return filepath.Dir(file), filepath.Base(file)
}

// Caller :
func Caller(fullpath bool) string {
	_, _, fn := trackCaller(1)
	if fullpath {
		return fn
	}
	return RmHeadToLast(fn, ".")
}

// ParentCaller :
func ParentCaller(fullpath bool) string {
	_, _, fn := trackCaller(2)
	if fullpath {
		return fn
	}
	return RmHeadToLast(fn, ".")
}

// FuncTrack : full path of func name
func FuncTrack(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
