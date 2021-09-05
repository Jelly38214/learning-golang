package main

import (
	"fmt"
	"path/filepath"
)

// .Ext, .Base, .Dir并不会去判断给定的路径是否真实存在
func main() {
	ext := filepath.Ext("./main.go")
	fmt.Println("extention: ", ext)

	filename := filepath.Base("./main.go")
	fmt.Println("filename: ", filename)

	directory := filepath.Dir("./main.go")
	fmt.Println("directory:", directory)
}
