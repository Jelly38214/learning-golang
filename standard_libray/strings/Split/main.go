package main

import (
	"fmt"
	"strings"
)

// 使用sep字符串去切割字符串s,最后返回一个切割后的字符串切片
func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a,b,c", "")) // ["a", ",", "b", ",", "c"]
}
