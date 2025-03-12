package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func SetupLogger() {
	Logger = log.New(os.Stdout, "banter: ", log.LstdFlags|log.Lshortfile)
}
