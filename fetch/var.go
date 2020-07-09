package fetch

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cdutwhu/debog/base"
	"github.com/cdutwhu/debog/fn"
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

	exist          = base.Exist
	rmTailFromLast = base.RmTailFromLast
	mustWriteFile  = base.MustWriteFile
	rmHeadToLast   = base.RmHeadToLast
	failOnErr      = fn.FailOnErr
	failPOnErr     = fn.FailPOnErr
	failPOnErrWhen = fn.FailPOnErrWhen
)

const tmFmt = "2006/01/02 15:04:05"
