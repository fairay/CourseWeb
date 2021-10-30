package utils

import (
	"log"
	"net/http"
	"os"
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

	Logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

func CloseLogger() {
	logFile.Close()
}


var LogHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Printf("%s REQUEST\t URL:%s \tAddress: %s", r.Method, r.URL, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}) 
}
