package main

import "fmt"

// 切片有3个字段: 指向底层数组的指针,切片访问的元素的个数(长度),切片允许增长到的元素个数(容量)

// 能否提前知道切片需要的容量通常会决定要如何创建切片

func main() {
	// 1.使用make创建并初始化切片, 其长度和容量都是5
	slice := make([]string, 5)
	fmt.Println(slice)

	// 分别指定长度和容量, 其长度为3,容量为5
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)

	// 2.使用字面量来创建并初始化, 长度和容量都是字面量元素个数
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice3)

	slice4 := []int{10, 20, 30}
	fmt.Println(slice4)

	// 3.使用索引声明切片
	slice5 := []string{99: "A"} // 长度和容量都是100的切片, 索引位置99的元素值为A
	fmt.Println(slice5)

	// 4.声明一个值为nil的切片(nil切片).只要在声明时不做任何初始化,就会创建一个nil切片
	var slice6 []int // nil切片, 指向底层数组的指针为nil, 长度为0, 容量为0
	fmt.Println("nil 切片: ", slice6)

	// 5.声明一个空切片, 底层数组包含0个元素, 长度和容量都为0; 表示空集合时很有用
	slice7 := make([]int, 0)
	slice8 := []int{}
	fmt.Println("空切片: ", slice7, slice8)

}
