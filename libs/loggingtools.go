package libs

import (
	"os"
	"time"
	"log"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)


func Init() {
	var acdate string = time.Now().Format("01-02-2006")
	var file, err = os.OpenFile("Logs/wbas_log" + acdate + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
        log.Fatal(err)
    }
 	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(logtext string) {
	Init()
	InfoLogger.Println(logtext)
}