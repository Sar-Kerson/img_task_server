package logger

import (
	"log"
	"os"
)

var (
	appLog *os.File
)

func init() {
	var err error
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	appLog, err = os.OpenFile("/tmp/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(appLog)
}

func Close() {
	appLog.Close()
}
