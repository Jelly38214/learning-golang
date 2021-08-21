package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer值, 并将一个字符串写入Buffer
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用Fprintf来将一个字符串拼接到Buffer里
	// 将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&b, "World") // Fprintf根据格式化说明符来格式化写入内容并输出到w,返回写入的字节数以及遇到的错误

	// 将Buffer的内容输出到标准输出设备
	// 将os.File值的地址作为io.Writer类型值传入
	b.WriteTo(os.Stdout)
}
