package main

import (
	"fmt"
	"strings"
)

// Index查找子串第一次出现的位置, 不存在则返回-1
// LastIndex则是返回最后一个匹配到位置
func main() {
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1
}
