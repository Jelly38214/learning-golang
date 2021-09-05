package main

import (
	"fmt"
	"strings"
)

// 判断字符串s包含字串的个数
func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("cheese", "ee"))
	fmt.Println(strings.Count("cheese", ""))  // 7, 整体看做一次匹配
	fmt.Println(strings.Count("cheese", "E")) // 0
	fmt.Println(strings.Count("chsesese", "se"))
}
