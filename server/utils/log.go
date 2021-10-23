package utils

import (
	"os"
	"log"
)

var (
	Logger *log.Logger
	logFile *os.File
)

func InitLogger() {
	logFile, err := os.OpenFile(Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(logFile, "DEBUG\t", log.Ldate|log.Ltime)
}

func CloseLogger() {
	logFile.Close()
}