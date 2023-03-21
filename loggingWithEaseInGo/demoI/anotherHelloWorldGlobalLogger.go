package demoI

import "log"

func AnotherHelloWorldWithGlobalLogger() {
	log.Println("This is from the file anotherHelloWorldGlobalLogger.go")
	CloseFile()
}
