package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("./demo.txt")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return
		}

		if os.IsPermission(err) {
			fmt.Println("没有权限对文件进行操作")
			return
		}

		fmt.Println("其他错误")
		return
	}

	fmt.Print("文件存在")
}
