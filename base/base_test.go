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

func TestMustWriteFile(t *testing.T) {
	str := "hello write"
	MustWriteFile("../a/b.log", []byte(str))
	MustWriteFile("/root/a/b.log", []byte(str))
}

func TestMustAppendFile(t *testing.T) {
	str := "hello append"
	MustAppendFile("../a/b.log", []byte(str), true)
	MustAppendFile("/root/a/b.log", []byte(str), true)
}

func TestCaller(t *testing.T) {
	fPln(Caller(false))
	fPln(Caller(true))
}

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

func TestColorPrint(t *testing.T) {
	// effect in bash
	fPln(B("hello"))
}
