package main

import "github.com/KianYang-Lee/goblog/loggingWithEaseInGo/demoI"

func main() {
	demoI.HelloWorld()
	demoI.HelloWorldWithLog()
	demoI.HelloWorldWithConfiguredLog()
	demoI.AnotherHelloWorldWithLog()
	demoI.HelloWorldWithGlobalLogger()
	demoI.AnotherHelloWorldWithGlobalLogger()
}
