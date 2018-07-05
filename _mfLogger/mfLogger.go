package _mfLogger

import (
	"log"
	"os"
	"sync"
)

type mfLogger struct {
	*log.Logger
	fileName string
}

var mfL *mfLogger
var once sync.Once

// GetInstance: Create a singleton instance of the mf logger
func GetInstance() *mfLogger {
	once.Do(func() {
		mfL = createLogger("mflogger.log")
	})
	return mfL
}

func createLogger(fname string) *mfLogger { // Create a logger instance
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	return &mfLogger{
		Logger:   log.New(file, "MetaFries ", log.Lshortfile),
		fileName: fname,
	}
}
