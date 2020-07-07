package fetch

import (
	"fmt"
	"strconv"
	"strings"
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
)

const tmFmt = "2006/01/02 15:04:05"
