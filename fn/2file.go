package fn

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cdutwhu/debog/base"
)

// FlagFileLog : is logging into local file
func FlagFileLog() bool {
	return log2file
}

// SetLog :
func SetLog(logpath string) {
	zone, offset := time.Now().Zone()
	if sHasSuffix(logpath, ".log") {
		logpath = logpath[:len(logpath)-4]
	}
	cat := "+"
	if offset < 0 {
		cat = "-"
	}

	logpath += fSf("@%s%s%4.1f%s", zone, cat, float32(offset/3600.0), ".log")
	if abspath, err := filepath.Abs(logpath); err == nil {
		base.MustAppendFile(abspath, nil, false)
		if f, err := os.OpenFile(abspath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
			mPathFile[abspath] = f
			log.SetFlags(log.LstdFlags) // log.SetFlags(log.LstdFlags | log.LUTC)
			log.SetOutput(f)
			log2file = true
		}
	}
}

// ResetLog : call once at the exit
func ResetLog() error {
	for logPath, f := range mPathFile {
		// delete empty error log
		fi, err := f.Stat()
		// FailOnErr("%v", err)
		if err != nil {
			return err
		}

		if fi.Size() == 0 {
			// FailOnErr("%v", os.Remove(logPath))
			if err := os.Remove(logPath); err != nil {
				return err
			}
		}
		// close
		f.Close()
	}
	mPathFile = make(map[string]*os.File)
	log.SetOutput(os.Stdout)
	log2file = false
	return nil
}
