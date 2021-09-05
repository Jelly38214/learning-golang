package main

import (
	"fmt"
	"strings"
)

// 将一系列字符串切片链接为一个字符串,字符串之间用sep来分隔
// Join就是Split的反向操作
func main() {
	s := []string{"foo", "bar", "baz"}

	fmt.Println(strings.Join(s, ", "))
}
