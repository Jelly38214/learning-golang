package main

import "fmt"

// 迭代map不能保证键值对返回的顺序,因为它使用了散列表

func main() {
	// 使用make创建键的类型是string, 值的类型是int的map
	dict := make(map[string]int)
	fmt.Println("dict: ", dict)

	// 字面量创建并初始化map
	dict1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println("dict1: ", dict1)

	// 创建一个空map,并赋值
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	fmt.Println("colors: ", colors)

	// 创建nil map,它不能用于存储健值对,会报错
	var nilMap map[string]string
	nilMap["Red"] = "#da1337"
}
