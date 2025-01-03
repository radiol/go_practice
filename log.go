package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!")
	log.Printf("%T %v", "test", "test")
	// log.Fatalln("error!")
	loggingToFile()
}

func LoggingSetting(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func loggingToFile() {
	LoggingSetting("test.log")
	_, err := os.Open("notExist.file")
	log.Fatalln(err)
}
