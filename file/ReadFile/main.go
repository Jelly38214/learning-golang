package main

import (
	"fmt"
	"io/ioutil"
)

// ReadFile 读取整个文件,不会将读取返回的EOF视为错误
func main() {
	content, err := ioutil.ReadFile("./README.md") // 路径相对于,go run命令执行所在目录路径

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
