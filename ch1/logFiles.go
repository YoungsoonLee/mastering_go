package ch1

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	pgname := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, pgname)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Logging in GO!")

}
