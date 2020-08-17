package fn

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cdutwhu/debog/base"
	clr "github.com/gookit/color"
)

var (
	fPt          = fmt.Print
	fPln         = fmt.Println
	fSf          = fmt.Sprintf
	fEf          = fmt.Errorf
	sHasSuffix   = strings.HasSuffix
	sLastIndex   = strings.LastIndex
	sReplaceAll  = strings.ReplaceAll
	sJoin        = strings.Join
	sSplit       = strings.Split
	sContains    = strings.Contains
	scParseFloat = strconv.ParseFloat

	mustAppendFile = base.MustAppendFile
	trackCaller    = base.TrackCaller
	caller         = base.Caller
	rmTailFromLast = base.RmTailFromLast

	R = clr.FgRed.Render
	G = clr.FgGreen.Render
	B = clr.FgBlue.Render
	Y = clr.FgYellow.Render
	W = clr.FgWhite.Render
)

const (
	tmFmt        = "2006/01/02 15:04:05 " // end with " " same as log.Printf
	logfile4test = "../a/b.log"
)

var (
	log2C          = true
	log2F          = false
	warnWithDetail = true

	mPathFile map[string]*os.File = make(map[string]*os.File)
	mFnType   map[string]string   = map[string]string{
		"Logger":    "INFO",
		"debug":     "DEBUG",
		"warnOnErr": "WARN",
		"failOnErr": "FAIL",
	}
	tmstr = func() string {
		return time.Now().Format(tmFmt)
	}
)
