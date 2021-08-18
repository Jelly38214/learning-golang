package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 创建一个定制的日志记录器, 需要创建一个Logger类型值

var (
	// Trace 记录所有的日志
	Trace *log.Logger
	// Info 重要的信息
	Info *log.Logger
	// Warning 需要注意的信息
	Warning *log.Logger
	// Error 非常严重的问题
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile) // 丢弃所有的输出

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")

	Info.Println("Special Infomation")

	Warning.Println("There is something you need to know about")

	Error.Println("Something has failed")
}
