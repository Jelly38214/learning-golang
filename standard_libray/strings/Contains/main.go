package main

import (
	"fmt"
	"strings"
)

// Contains: 判断字符串s是否包含子串substr
func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}
