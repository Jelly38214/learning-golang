package main

import (
	"fmt"
	"strings"
)

// 将字符串前后位置的指定字符去掉
func main() {
	fmt.Println(strings.Trim("  www.google.com  ", " ")) // 去掉前后空字符串
	fmt.Println(strings.Trim("  www.google.com  ", "-")) // 去掉前后字符-,本次操作无影响,因为字符串前后没有-

	// 将s前后端素有空白都去掉
	fmt.Println(strings.TrimSpace("  www.google.com  ")) // 去掉前后空字符串
}
