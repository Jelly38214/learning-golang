package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println写到标准日志记录器
	log.Println("message")

	// Fatalln在调用Println()之后会接着调用os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln在调用Println()之后会接着调用panic()
	log.Panicln("panic message")
}
