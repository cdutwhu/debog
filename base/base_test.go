package base

import "testing"

func TestExist(t *testing.T) {
	fPln(Exist(1, 1))
	fPln(Exist(1))
	fPln(Exist(1, 2, 3))
	fPln(Exist(1, 3, 2, 1))
}

func TestRmTailFromLast(t *testing.T) {
	fPln(RmTailFromLast("abc..def..ghi", ".."))
}

func TestRmHeadToLast(t *testing.T) {
	fPln(RmHeadToLast("abc..def..ghi", ".."))
}

// echo cppcli | sudo -S env "PATH=$PATH" go test -v ./ -run TestMustWriteFile
// echo cppcli | sudo -S rm -rf /root/a/

func TestMustWriteFile(t *testing.T) {
	str := "hello write"
	MustWriteFile("../a/b/c.log", []byte(str))
	// MustWriteFile("/root/a/b.log", []byte(str))
}

func TestMustAppendFile(t *testing.T) {
	str := "hello append"
	MustAppendFile("../a/b/c.log", []byte(str), true)
	// MustAppendFile("/root/a/b.log", []byte(str), true)
}

// ------------------------------------------ //

type Tp struct {
}

func (t *Tp) hello() {
	fPln("hello")
}

func hello1() {
	fPln("hello")
}

func TestFuncTrack(t *testing.T) {
	p := Tp{}
	fPln(FuncTrack(p.hello))
	fPln(FuncTrack(hello1))
}

func TestTrackCallerIn(t *testing.T) {
	fPln(trackCaller(0))
}

func TestTrackCaller(t *testing.T) {
	fPln(TrackCaller(0))
}

func TestCallerSrc(t *testing.T) {
	fPln(CallerSrc())
}

func TestCaller(t *testing.T) {
	fPln(Caller(false))
	fPln(Caller(true))
}

func Wrapper() {
	fPln(ParentCaller(false))
	fPln(ParentCaller(true))
}

func TestParentCaller(t *testing.T) {
	Wrapper()
}
