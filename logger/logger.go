package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

func init() {
	dir := "logger"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.Create(filepath.Join(dir, filepath.Base("general-log.txt")))
	if err != nil {
		fmt.Println("error while create directory for log")
		os.Exit(1)
	}
	GeneralLogger = log.New(f, "General Logger:\t", log.Lshortfile)
	ErrorLogger = log.New(f, "Error Logger:\t", log.Lshortfile)
}
