package libs

import (
	"os"
	"time"
	"log"
	"fmt"
	)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)


func Init() {

	if _, err := os.Stat("Logs"); os.IsNotExist(err) {
		err := os.Mkdir("Logs", 0755)
		if err != nil {
			fmt.Println("Fatal Error: Could not create Logs Folder, Server Terminating...\n Exception" + err.Error())
		 } else {
		 	fmt.Print("Successfully created Logs Folder...")
		 } 
	}
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

func LogError(logtext string) {
	Init()
	ErrorLogger.Println(logtext)
}

func LogWarning(logtext string) {
	Init()
	WarningLogger.Println(logtext)
}
