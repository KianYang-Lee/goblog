package demoI

import (
	"log"
	"os"
)

func HelloWorldWithConfiguredLog() {
	f, _ := os.OpenFile("loggingWithEaseInGo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // Ignore error from opening file
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags|log.Lshortfile)
	logger.Println("Hello, World!")
}
