package fn

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cdutwhu/debog/base"
)

var (
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
	blue           = base.B
	red            = base.R
	green          = base.G
	yellow         = base.Y
	white          = base.W
	decolor        = base.DeColor
)

const (
	tmFmt        = "2006/01/02 15:04:05 " // end with " " same as log.Printf
	logfile4test = "../a/b.log"
)

var (
	log2file                      = false
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
