package demoI

import "log"

func AnotherHelloWorldWithLog() {
	// Did not perform configuration step
	log.Println("This is log written by another file")

	// Terminal
	// $ go run ...
	// 2023/03/20 22:29:05 This is log written by another file
}
