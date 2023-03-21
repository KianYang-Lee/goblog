package demoI

import (
	"log"
	"os"
)

var f *os.File

func HelloWorldWithGlobalLogger() {
	f, _ := os.OpenFile("loggingWithEaseInGo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // Ignore error from opening file
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("Global Logger : ")
	log.Println("Hello, World!")
}

func CloseFile() {
	f.Close()
}
