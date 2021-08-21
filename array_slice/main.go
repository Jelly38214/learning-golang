package main

import "fmt"

// 数组是一个长度固定的数据类型,用于存储一段具有相同的类型的元素的连续块

func main() {

	// 声明数组时需要指定内部存储的数据的类型,以及需要存储的元素的数量(数组长度)
	// 一旦声明,数组长度不可变.如果需要存储更多的元素,就需要先创建一个更长的数组,再把原来数组里的值复制到新数组里
	var array [5]int              // 数组初始化时,数组内每个元素都初始化为对应类型的零值
	fmt.Println("array: ", array) // [0 0 0 0 0]

	// 快速创建数组并初始化的方式是使用数组字面量, 数组字面量允许声明数组里元素的数量同时指定每个元素的值
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2: ", array2)

	// 使用...替代数组的长度, Go语言会根据初始化时数组元素的数量来确定该数组的长度
	array3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("array3: ", array3)

	// 声明一个具体长度,并用具体值初始化索引为1和2元素, 其余索引位置上的元素使用类型零值
	array4 := [5]int{1: 10, 2: 20}
	fmt.Println("array4: ", array4)

	// 声明一个元素全是指针的数组
	arrayWithPointer := [5]*int{0: new(int), 1: new(int)}
	fmt.Printf("arrayWithPointer: %v\n", arrayWithPointer)

	// 使用*运算符就可以访问元素指针指向的值
	*arrayWithPointer[0] = 10
	*arrayWithPointer[1] = 20
	fmt.Printf("arrayWithPointer: %v\n", arrayWithPointer)

	// 类型相同的数组指的是数组长度和元素类型一致, 它们之间可以互相赋值
	array = array2 // 将array2的值复制到array
	fmt.Println("array: ", array)

}
