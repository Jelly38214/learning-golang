package main

import (
	"fmt"
	"strings"
)

// 将s中前n个不重叠old字符串都替换成new字串, 如果n<0则会替换所有old字串
func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // 第三个k不会被替换成ky
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //  全替换
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))  //  全替换
}
